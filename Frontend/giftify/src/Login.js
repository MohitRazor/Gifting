import React, { useState } from 'react';
import './Login.css';

const Login = () => {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const [errors, setErrors] = useState({});
  const [successMessage, setSuccessMessage] = useState('');

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const validateForm = () => {
    const { email, password } = formData;
    const newErrors = {};

    if (!email) newErrors.email = 'Email is required';
    else if (!/\S+@\S+\.\S+/.test(email)) newErrors.email = 'Invalid email format';
    if (!password) newErrors.password = 'Password is required';
    else if (password.length < 6) newErrors.password = 'Password must be at least 6 characters';
    return newErrors;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const validationErrors = validateForm();
    if (Object.keys(validationErrors).length > 0) {
      setErrors(validationErrors);
      setSuccessMessage('');
    } else {
      setErrors({});

      try {
        const response = await fetch('http://localhost:3001/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: formData.email,
            password: formData.password,
          }),
        });

        if (response.ok) {
          setSuccessMessage('Login successful!');
          setFormData({
            email: '',
            password: '',
          });
        } else {
          const errorData = await response.json();
          console.error('Login failed:', errorData);
          setErrors({ backendError: errorData.error });
          setSuccessMessage('');
        }
      } catch (error) {
        console.error('Error during login:', error);
        setErrors({ networkError: 'Failed to connect to server' });
        setSuccessMessage('');
      }
    }
  };

  return (
    <div className="container">
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="email" style={{ color: 'white' }}>Email</label>
          <input
            type="email"
            id="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
          />
          {errors.email && <span className="error" style={{ color: 'red', fontSize: '14px', marginTop: '5px' }}>{errors.email}</span>}
        </div>
        <div className="form-group">
          <label htmlFor="password" style={{ color: 'white' }}>Password</label>
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
          />
          {errors.password && <span className="error" style={{ color: 'red', fontSize: '14px', marginTop: '5px' }}>{errors.password}</span>}
        </div>
        <button type="submit" style={{ color: 'white', background: '#ff666f', padding: '10px', marginTop: '10px', border: 'none', borderRadius: '18px' }}>Login</button>
      </form>
      {successMessage && <div style={{ color: 'white', padding: '10px', marginTop: '10px' }}>{successMessage}</div>}
      {errors.backendError && <div style={{ color: 'red', padding: '10px', marginTop: '10px' }}>{errors.backendError}</div>}
      {errors.networkError && <div style={{ color: 'red', padding: '10px', marginTop: '10px' }}>{errors.networkError}</div>}
    </div>
  );
};

export default Login;
