import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import '../styles/Animals.css';
import axios from "axios";


interface Animal {
  id: number;
  name: string;
  species: string;
  age: number;
}


const Animals: React.FC = () => {
  const [animals, setAnimals] = useState<Animal[]>([]);
  const [searchTerm, setSearchTerm] = useState<string>("");
  const navigate = useNavigate();


  const fetchAnimals = async () => {
    try { // TODO: sep api env
      const response = await axios.get("http://localhost:8080/api/v1/animals");
      const data = response.data;

      if (Array.isArray(data)) {
        setAnimals(data as Animal[]);
      } else {
        console.error("Invalid data format", data);
      }
    } catch (error) {
      console.error("Failed to fetch animals", error);
    }
  };

  useEffect(() => {
    fetchAnimals();
  }, []);

  const handleSearch = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const handleAnimalClick = (id: number) => {
    navigate(`/animal/${id}`);
  };

  return (
    <div className="container">
      <h1>Animals</h1>
      <input
        type="text"
        value={searchTerm}
        onChange={handleSearch}
        placeholder="Search for an animal..."
      />
      <ul className="animal-list">
        {animals
          .filter(animal => animal.name.toLowerCase().includes(searchTerm.toLowerCase()))
          .map(animal => (
            <li key={animal.id} className="animal-item">
              <span className="animal-name">{animal.name}</span>
              <button className="animal-button" onClick={() => handleAnimalClick(animal.id)}>
                View Details
              </button>
            </li>
          ))}
      </ul>
    </div>
  );
};


export default Animals;
