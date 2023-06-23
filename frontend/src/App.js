/*
 * File: App.js
 * File Created: Monday, 12th June 2023 12:36:14 am
 * Last Modified: Monday, 12th June 2023 1:57:11 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'; // Update import statement
import { createGlobalStyle } from 'styled-components';
import LicenseManagement from './components/LicenseManagement';
import CourseManagement from './components/CourseManagement';
import Navbar from './components/Navbar';

const GlobalStyle = createGlobalStyle`
  body {
    background-color: #e0e5ec;
    font-family: Arial, sans-serif;
  }
`;

function App() {
  return (
    <Router>
      <GlobalStyle />
      <Navbar />
      <Routes> {/* Updated component name */}
        <Route path="/licenses" element={<LicenseManagement />} />
        <Route path="/courses" element={<CourseManagement />} />
      </Routes> {/* Updated component name */}
    </Router>
  );
}

export default App;
