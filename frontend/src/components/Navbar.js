/*
 * File: Navbar.js
 * File Created: Monday, 12th June 2023 12:41:23 am
 * Last Modified: Monday, 12th June 2023 12:41:35 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';

const Nav = styled.nav`
  background-color: #e0e5ec;
  display: flex;
  justify-content: space-around;
  padding: 1em;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
`;

const NavLink = styled(Link)`
  color: #333;
  text-decoration: none;
`;

function Navbar() {
  return (
    <Nav>
      <NavLink to="/licenses">License Management</NavLink>
      <NavLink to="/courses">Course Management</NavLink>
    </Nav>
  );
}

export default Navbar;
