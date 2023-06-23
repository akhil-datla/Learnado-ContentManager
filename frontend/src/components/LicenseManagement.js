/*
 * File: LicenseManagement.js
 * File Created: Monday, 12th June 2023 12:35:30 am
 * Last Modified: Tuesday, 13th June 2023 3:05:21 pm
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

import React, { useState } from 'react';
import axios from 'axios';
import styled from 'styled-components';
import { AiOutlineCopy } from 'react-icons/ai';

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2em;
  grid-template-columns: repeat(2, 1fr);
  grid-gap: 20px;
  margin-top: 20px;
`;

const Form = styled.form`
  background: #e0e5ec;
  width: 300px;
  height: 200px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 1em;
  border-radius: 0.5em;
  box-shadow:  20px 20px 60px #bec3c9, 
             -20px -20px 60px #ffffff;
`;

const Input = styled.input`
  width: 80%;
  padding: 0.5em;
  margin: 0.5em 0;
  border: none;
  border-radius: 0.5em;
  box-shadow: inset 8px 8px 8px #cbced1,
              inset -8px -8px 8px #ffffff;
`;

const Button = styled.button`
  padding: 0.5em 1em;
  border: none;
  border-radius: 0.5em;
  box-shadow:  3px 3px 5px #b8b9be, 
             -3px -3px 5px #ffffff;
  cursor: pointer;
`;

const IconButton = styled(Button)`
  display: flex;
  align-items: center;
  justify-content: center;
`;

const LicenseList = styled.ul`
  max-height: 200px;
  overflow-y: scroll;
  padding: 0.5em;
  border-radius: 0.5em;
  box-shadow: inset 8px 8px 8px #cbced1,
              inset -8px -8px 8px #ffffff;
  list-style: none;
  margin-bottom: 1em;
`;

function LicenseManagement() {
  const [courseId, setCourseId] = useState('');
  const [numLicenses, setNumLicenses] = useState(1);
  const [licenseKeys, setLicenseKeys] = useState([]);
  const [licenseKey, setLicenseKey] = useState('');

  const generateLicenses = async e => {
    e.preventDefault();
    
    try {
      const res = await axios.post('http://localhost:8080/licenses/create', { courseID: courseId, num: numLicenses.toString() });
      setLicenseKeys(res.data.licenseKeys);
    } catch(err) {
      console.error(err);
      alert('An error occurred.');
    }
  };  

  const copyToClipboard = () => {
    const tempInput = document.createElement("textarea");
    tempInput.style = "position: absolute; left: -1000px; top: -1000px";
    tempInput.value = licenseKeys.join('\n');
    document.body.appendChild(tempInput);
    tempInput.select();
    document.execCommand("copy");
    document.body.removeChild(tempInput);
    alert('Licenses copied to clipboard!');
  };
  

  const revokeLicense = async e => {
    e.preventDefault();
    
    try {
      await axios({
        method: 'delete',
        url: 'http://localhost:8080/licenses/revoke',
        data: { licenseKey },
      });
      alert('License revoked');
    } catch(err) {
      console.error(err);
      alert('An error occurred.');
    }
  };

  return (
    <Container>
      <Form onSubmit={generateLicenses}>
        <Input 
          type="text" 
          value={courseId} 
          onChange={e => setCourseId(e.target.value)} 
          placeholder="Course ID" 
          required 
        />
        <Input 
          type="number" 
          value={numLicenses} 
          onChange={e => setNumLicenses(e.target.value)} 
          placeholder="Number of Licenses" 
          required 
        />
        <Button type="submit">Generate Licenses</Button>
      </Form>
      
      {licenseKeys.length > 0 && 
        <div>
          <h3>Generated Licenses:</h3>
          <LicenseList>
            {licenseKeys.map((licenseKey, index) => <li key={index}>{licenseKey}</li>)}
          </LicenseList>
          <IconButton onClick={copyToClipboard}>
            <AiOutlineCopy size={20} />
          </IconButton>
        </div>
      }
      
      <Form onSubmit={revokeLicense}>
        <Input 
          type="text" 
          value={licenseKey} 
          onChange={e => setLicenseKey(e.target.value)} 
          placeholder="License Key" 
          required 
        />
        <Button type="submit">Revoke License</Button>
      </Form>
    </Container>
  );
}

export default LicenseManagement;