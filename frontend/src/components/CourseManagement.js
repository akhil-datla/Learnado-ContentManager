/*
 * File: CourseManagement.js
 * File Created: Monday, 12th June 2023 12:35:55 am
 * Last Modified: Tuesday, 13th June 2023 3:00:12 pm
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styled from 'styled-components';

const Container = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  padding: 1em;
`;

const FormContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2em;
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

const CourseList = styled.ul`
  width: 300px;
  max-height: 500px;  // adjust this as per your requirement
  overflow-y: auto;  // this will add a scroll bar when content exceeds max-height
  list-style-type: none;
  padding: 1em;
  border-radius: 0.5em;
  box-shadow:  20px 20px 60px #bec3c9, 
             -20px -20px 60px #ffffff;
`;

const CourseListItem = styled.li`
  background: #e0e5ec;
  padding: 1em;
  border-radius: 0.5em;
  box-shadow:  20px 20px 60px #bec3c9, 
             -20px -20px 60px #ffffff;
`;

function CourseManagement() {
  const [courseId, setCourseId] = useState('');
  const [courseName, setCourseName] = useState('');
  const [filePath, setFilePath] = useState('');
  const [courses, setCourses] = useState([]);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const res = await axios.get('http://localhost:8080/courses/all');
        setCourses(res.data);
      } catch (err) {
        console.error(err);
        alert('An error occurred while fetching courses.');
      }
    };
    fetchCourses();
  }, []);

  const createCourse = async e => {
    e.preventDefault();
    
    try {
      const res = await axios.post('http://localhost:8080/courses/create', { name: courseName, filepath: filePath });
      alert(`Course ID: ${res.data.courseID}`);
    } catch(err) {
      console.error(err);
      alert('An error occurred.');
    }
  };

  const updateCourse = async e => {
    e.preventDefault();
    
    try {
      await axios.post('http://localhost:8080/courses/update', { id: courseId, name: courseName, filepath: filePath });
      alert('Course updated');
    } catch(err) {
      console.error(err);
      alert('An error occurred.');
    }
  };

  const deleteCourse = async e => {
    e.preventDefault();
    
    try {
      await axios.delete('http://localhost:8080/courses/delete', { data: { id: courseId }});
      alert('Course deleted');
    } catch(err) {
      console.error(err);
      alert('An error occurred.');
    }
  };

  return (
    <Container>
      <CourseList>
        {courses.map(course => (
          <CourseListItem key={course.ID}>
            <p><strong>ID:</strong> {course.ID}</p>
            <p><strong>Name:</strong> {course.Name}</p>
            <p><strong>File Path:</strong> {course.Filepath}</p>
          </CourseListItem>
        ))}
      </CourseList>

      <FormContainer>
        <Form onSubmit={createCourse}>
          <Input 
            type="text" 
            value={courseName} 
            onChange={e => setCourseName(e.target.value)} 
            placeholder="Course Name" 
            required 
          />
          <Input 
            type="text" 
            value={filePath} 
            onChange={e => setFilePath(e.target.value)} 
            placeholder="File Path" 
            required 
          />
          <Button type="submit">Create Course</Button>
        </Form>

        <Form onSubmit={updateCourse}>
          <Input 
            type="text" 
            value={courseId} 
            onChange={e => setCourseId(e.target.value)} 
            placeholder="Course ID" 
            required 
          />
          <Input 
            type="text" 
            value={courseName} 
            onChange={e => setCourseName(e.target.value)} 
            placeholder="Course Name" 
            required 
          />
          <Input 
            type="text" 
            value={filePath} 
            onChange={e => setFilePath(e.target.value)} 
            placeholder="File Path" 
            required 
          />
          <Button type="submit">Update Course</Button>
        </Form>

        <Form onSubmit={deleteCourse}>
          <Input 
            type="text" 
            value={courseId} 
            onChange={e => setCourseId(e.target.value)} 
            placeholder="Course ID" 
            required 
          />
          <Button type="submit">Delete Course</Button>
        </Form>
      </FormContainer>
    </Container>
  );
}

export default CourseManagement;
