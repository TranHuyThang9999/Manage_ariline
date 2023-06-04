import React, { useState } from 'react';
import { Button, Form, Input, DatePicker, InputNumber } from 'antd';
import moment from 'moment';
import axios from 'axios';
import './Form.css';

const CreateUserForm = () => {
  const [loading, setLoading] = useState(false);

  const onFinish = async (values) => {
    try {
      setLoading(true);
      values.amount = parseInt(values.amount);

      const response = await axios.post('http://localhost:8080/user/create', values);
      console.log(response.data);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="form-container">
      <Form onFinish={onFinish}>
        <Form.Item
          name="user_name"
          label="User Name"
          rules={[{ required: true, message: 'Please enter your user name' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="phone_number"
          label="Phone Number"
          rules={[{ required: true, message: 'Please enter your phone number' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="date_birth"
          label="Date of Birth"
          rules={[{ required: true, message: 'Please select your date of birth' }]}
        >
          <DatePicker format="DD/MM/YYYY" />
        </Form.Item>
        <Form.Item
          name="email"
          label="Email"
          rules={[{ required: true, message: 'Please enter your email' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="password"
          label="Password"
          rules={[{ required: true, message: 'Please enter your password' }]}
        >
          <Input.Password />
        </Form.Item>
        <Form.Item
          name="address"
          label="Address"
          rules={[{ required: true, message: 'Please enter your address' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="number_cmnd"
          label="ID Card Number"
          rules={[{ required: true, message: 'Please enter your ID card number' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="nationality"
          label="Nationality"
          rules={[{ required: true, message: 'Please enter your nationality' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="language"
          label="Language"
          rules={[{ required: true, message: 'Please enter your language' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="amount"
          label="Amount"
          rules={[{ required: true, message: 'Please enter the amount' }]}
        >
          <InputNumber/>
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit" loading={loading}>
          sign up
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default CreateUserForm;
