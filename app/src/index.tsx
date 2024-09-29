import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import AnimalDetails from './pages/AnimalDetails';
import { useParams } from 'react-router-dom';
import Animals from './pages/Animals';
import ReactDOM from 'react-dom';
import Auth from './pages/Auth';
import React from 'react';


const AnimalDetailsWrapper = () => {
  const { id } = useParams();
  return <AnimalDetails animalId={Number(id)} />;
};


const App = () => (
  <Router>
    <Routes>
      <Route path="/" element={<Auth />} />
      <Route path="/animals" element={<Animals />} />
      <Route path="/animal/:id" element={<AnimalDetailsWrapper />} />
    </Routes>
  </Router>
);


ReactDOM.render(<App />, document.getElementById('root'));
