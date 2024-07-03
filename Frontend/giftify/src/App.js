import React from 'react';
import Register from './Register';
import './App.css';

function App() {
  // URL for the slogan image
  const sloganImageUrl = 'https://cdn.prod.website-files.com/6595532c1f7a37b39fc32ae8/6595532c1f7a37b39fc32b3f_card-1.svg';

  return (
    <div className="App">
      <header className='nav'>
        <div className="search-bar">
          <input type="text" placeholder="Search..." />
          <button>Search</button>
        </div>
        <div className="auth-buttons">
          <button>Login</button>
          <button>Signup</button>
        </div>
      </header>
      <section className='content-section'>
        <div className='content-body'>
        <img src={sloganImageUrl} className="App-slogan" alt="slogan" />
        <p className='slogan-text'>
          Gifts for needs and happiness
        </p>
        </div>
        <div id='register' className='register'>
        <Register />
        </div>
      </section>
      {/* <header className="App-header">
        <div className="header-content">
        <div className="search-bar">
            <input type="text" placeholder="Search..." />
            <button>Search</button>
          </div>
          <img src={sloganImageUrl} className="App-slogan" alt="slogan" />
          <p>
            Gifts for needs and happiness
          </p>
          <div className="auth-buttons">
            <button>Login</button>
            <button>Signup</button>
          </div>
        </div>
      </header> */}
      {/* <main>
        <Register />
      </main> */}
    </div>
  );
}

export default App;