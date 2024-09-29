import React, { useState, useEffect } from "react";
import axios from "axios";


const Search: React.FC = () => {
  const [searchResults, setSearchResults] = useState<string[]>([]);
  const [query, setQuery] = useState("");


  const handleSearch = async () => {
    try {
      const response = await axios.get(`http://localhost:8000/autocomplete?q=${query}`);
      const data = response.data;

      if (Array.isArray(data)) {
        setSearchResults(data);
      } else {
        console.error("Unexpected data format", data);
      }
    } catch (error) {
      console.error("Search request failed", error);
    }
  };

  useEffect(() => {
    if (query) {
      handleSearch();
    }
  }, [query]);

  return (
    <div>
      <input
        type="text"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        placeholder="Search for an animal"
      />
      <ul>
        {searchResults.map((result, index) => (
          <li key={index}>{result}</li>
        ))}
      </ul>
    </div>
  );
};


export default Search;
