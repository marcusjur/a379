import { useParams, useNavigate } from "react-router-dom";
import React, { useEffect, useState } from "react";
import '../styles/AnimalDetails.css';
import axios from "axios";


interface AnimalDetailsProps {
  animalId: number;
}


const AnimalDetails: React.FC<AnimalDetailsProps> = ({ animalId }) => {
  const { id } = useParams<{ id: string }>();
  const [animal, setAnimal] = useState<any>(null);
  const navigate = useNavigate();


  const fetchAnimalDetails = async () => {
    try { // TODO: sep api env
      const response = await axios.get(`http://localhost:8080/api/v1/animal/${id}`);
      setAnimal(response.data);
    } catch (error) {
      console.error("Failed to fetch animal details", error);
    }
  };


  useEffect(() => {
    fetchAnimalDetails();
  }, [id]);


  if (!animal) return <div>Loading...</div>;


  const speciesTag = animal.tags.find((tag: string) => tag.includes(":species")) as string | undefined;
  const ageTag = animal.tags.find((tag: string) => tag.includes(":age")) as string | undefined;

  const species = speciesTag ? speciesTag.split(":")[0] : null;
  const age = ageTag ? ageTag.split(":")[0] : null;


  return (
    <div className="container">
      <h1>{animal.name}</h1>
      {species && <p>Species: {species}</p>}
      <p>Owner: {animal.owner}</p>
      {age && <p>Age: {age}</p>}
      {animal.tags.filter((tag: string) => !tag.includes(":species") && !tag.includes(":age")).length > 0 && (
        <p>Tags: {animal.tags.filter((tag: string) => !tag.includes(":species") && !tag.includes(":age")).join(", ")}</p>
      )}
      <img src={animal.image} alt={`${animal.name}`} />

      <button className="go-back" onClick={() => navigate(-1)}>
        Go Back
      </button>
    </div>
  );
};


export default AnimalDetails;
