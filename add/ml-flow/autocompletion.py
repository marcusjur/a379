from fastapi import FastAPI, Header, Body, HTTPException
from datetime import datetime
import transformers
import hashlib
import uvicorn
import ctypes
import json


def validate_session_key(json_str):
    """
    Validate the session key from the provided JSON string.
    """
    try:
        session_info = json.loads(json_str)
        session_key = session_info.get("Session")
        created_time = session_info.get("Created")

        if not session_key or not created_time:
            print("Missing session key or created time in JSON")
            return False

        expected_key = calculate_session_key(created_time)
        return session_key == expected_key

    except json.JSONDecodeError as e:
        print("Error parsing JSON:", e)
        return False


def calculate_session_key(created_time):
    """
    Helper function to generate the expected session key based on the current date and time.
    """
    current_date = datetime.now().strftime("%Y-%m-%d")
    noise_pattern = "some-fixed-noise-pattern"

    combined_string = current_date + created_time + noise_pattern
    print(f"Combined string: {combined_string}")
    sha256_hash = hashlib.sha256(combined_string.encode()).hexdigest()

    return sha256_hash

app = FastAPI()
tokenizer = transformers.AutoTokenizer.from_pretrained("babylm/babyllama-100m-2024")
model = transformers.pipeline("text-generation", model="babylm/babyllama-100m-2024",
    tokenizer=tokenizer)


@app.get("/api/v1/autocompletion")
async def autocompletion(session: str = Header(..., alias="Authorization"), text: dict = Body(...)):
    if not validate_session_key(session.encode('utf-8')):
        raise HTTPException(status_code=401, detail="Session invalid")

    result = model(text['text'], max_new_tokens=4, return_full_text=False)
    print(f"Text: {text['text']}")
    print(f"Result: {result}")
    return json.dumps(result[0])


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
