import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import '../styles/Animals.css';
import axios from "axios";

interface Animal {
  id: number;
  name: string;
  owner: string; // Added owner field
  tags: string; // Added tags field
  image: string; // Added image field
}

const Animals: React.FC = () => {
  const [animals, setAnimals] = useState<Animal[]>([]);
  const [searchTerm, setSearchTerm] = useState<string>("");
  const [newAnimal, setNewAnimal] = useState<Animal>({
    id: 0,
    name: "",
    owner: "",
    tags: "",
    image: ""
  });
  const [showModal, setShowModal] = useState<boolean>(false);
  const navigate = useNavigate();

  const fetchAnimals = async () => {
    try {
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

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = event.target;
    setNewAnimal(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      const response = await axios.post("http://localhost:8080/api/v1/animals", newAnimal);
      const addedAnimal = { ...newAnimal, id: (response.data as { id: number }).id };
      setAnimals(prevAnimals => [...prevAnimals, addedAnimal]);
      setNewAnimal({ id: 0, name: "", owner: "", tags: "", image: "" });
      setShowModal(false);
    } catch (error) {
      console.error("Failed to add animal", error);
    }
  };

  const handleModalToggle = () => {
    setShowModal(prev => !prev);
  };

  const handleDelete = async (id: number) => {
    const confirmDelete = window.confirm("Are you sure you want to delete this animal?");
    if (confirmDelete) {
      try {
        await axios.delete(`http://localhost:8080/api/v1/animal/${id}`);
        setAnimals(prevAnimals => prevAnimals.filter(animal => animal.id !== id)); // Remove animal from state
      } catch (error) {
        console.error("Failed to delete animal", error);
      }
    }
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
              <button className="delete-button" onClick={() => handleDelete(animal.id)}>
                Delete
              </button>
            </li>
          ))}
      </ul>

      <button className="add-animal-button" onClick={handleModalToggle}>
        Add Animal
      </button>

      {showModal && (
        <div className="modal">
          <div className="modal-content">
            <h2>Add Animal</h2>
            <form onSubmit={handleSubmit}>
              <input
                type="text"
                name="name"
                value={newAnimal.name}
                onChange={handleInputChange}
                placeholder="Name"
                required
              />
              <input
                type="text"
                name="owner"
                value={newAnimal.owner}
                onChange={handleInputChange}
                placeholder="Owner"
                required
              />
              <input
                type="text"
                name="tags"
                value={newAnimal.tags}
                onChange={handleInputChange}
                placeholder="Tags (comma-separated)"
                required
              />
              <input
                type="text"
                name="image"
                value={newAnimal.image}
                onChange={handleInputChange}
                placeholder="Image URL"
                required
              />
              <button type="submit">Add Animal</button>
              <button type="button" onClick={handleModalToggle}>Cancel</button>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};

export default Animals;
