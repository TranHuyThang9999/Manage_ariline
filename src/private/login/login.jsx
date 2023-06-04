import React, { useState } from 'react';
import axios from 'axios';
import { setAuthToken } from './axiosConfig';
import './style.css';
import NextPage2 from '../user/inforfilghr';
import CreateUserForm from '../create_user/create_user';

const LoginForm = () => {
  const [phoneNumber, setPhoneNumber] = useState('');
  const [password, setPassword] = useState('');
  const [phoneNumberError, setPhoneNumberError] = useState(false);
  const [passwordError, setPasswordError] = useState(false);
  const [loginSuccess, setLoginSuccess] = useState(false);
  const [showCreateUserForm, setShowCreateUserForm] = useState(false);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setPhoneNumberError(false);
    setPasswordError(false);

    if (phoneNumber.trim() === '') {
      setPhoneNumberError(true);
      return;
    }

    if (password.trim() === '') {
      setPasswordError(true);
      return;
    }

    try {
      const response = await axios.post(
        'http://localhost:8080/user/login',
        {
          phone_number: phoneNumber,
          password: password,
        },
        {
          headers: {
            'Content-Type': 'application/json',
          },
          withCredentials: true,
        }
      );

      if (response.status === 200) {
        alert('Đăng nhập thành công');
        const token = response.data.token;
        console.log('Token:', token);

        document.cookie = `token=${token}; expires=7`;
       document.cookie = `phone_number=${phoneNumber}`;

        localStorage.setItem('phone_number', phoneNumber);
        localStorage.setItem('token',token)
        setAuthToken(token);
        setLoginSuccess(true);
      } else {
        alert('Đăng nhập thất bại, vui lòng kiểm tra lại tài khoản');
        setPassword('');
      }
    } catch (error) {
      console.error('Lỗi trong quá trình đăng nhập:', error);
      alert('Đăng nhập thất bại, vui lòng kiểm tra lại tài khoản');
      setPassword('');
    }
  };

  const handleSignUp = (event) => {
    event.preventDefault();
    setShowCreateUserForm(true);
  };

  return (
    <div>
      {!loginSuccess && !showCreateUserForm ? (
        <form onSubmit={handleSubmit} className="login-form">
          <h2>Account Login</h2>
          <div>
            <input
              type="text"
              id="phone_number"
              value={phoneNumber}
              onChange={(event) => setPhoneNumber(event.target.value)}
              className="form-input"
              placeholder="Nhập số điện thoại"
            />
            {phoneNumberError && (
              <p style={{ color: 'red' }}>Số điện thoại không được để trống</p>
            )}
          </div>
          <div>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
              className="form-input"
              placeholder="Nhập mật khẩu"
            />
            {passwordError && (
              <p style={{ color: 'red' }}>Mật khẩu không được để trống</p>
            )}
          </div>
          <button type="submit" className="form-button">
            Login
          </button>
          <p>
            Not registered?{' '}
            <button className="signup-link" onClick={handleSignUp}>
              Sign up
            </button>
          </p>
        </form>
      ) : null}

      {showCreateUserForm && <CreateUserForm/>}

      {loginSuccess && <NextPage2 />}
    </div>
  );
};

export default LoginForm;
