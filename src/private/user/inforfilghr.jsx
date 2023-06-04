import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Table } from 'antd';
import Cookies from 'js-cookie';

const Page2 = () => {
  const [flightInfo, setFlightInfo] = useState([]);

  useEffect(() => {
    const token = Cookies.get('token');

    const getFlightInfo = async () => {
      try {
        const response = await axios.get('http://localhost:8080/user/info/flight', {
          headers: {
            Authorization: `${token}`,
          },
        });
        setFlightInfo(response.data.body); // Chỉnh sửa ở đây
      } catch (error) {
        console.error('Error while fetching flight information:', error);
      }
    };

    getFlightInfo();
  }, []);

  const columns = [
    {
      title: 'ID',
      dataIndex: 'flight_id',
      key: 'name_flight',
    },
    {
      title: 'Name Flight',
      dataIndex: 'name_flight',
      key: 'name_flight',
    },
    {
      title: 'Airline',
      dataIndex: 'name_airline',
      key: 'flight_id',
    },
    {
      title: 'Departure',
      dataIndex: 'departure',
      key: 'departure',
    },
    {
      title: 'Destination',
      dataIndex: 'destination',
      key: 'destination',
    },
    {
      title: 'Departure Time',
      dataIndex: 'departure_time',
      key: 'departure_time',
    },
    {
      title: 'Destination Time',
      dataIndex: 'destination_time',
      key: 'destination_time',
    },
    {
      title: 'Ticket Type',
      dataIndex: 'ticket_type',
      key: 'ticket_type',
    },
    {
      title: 'Ticket Fare',
      dataIndex: 'fare',
      key: 'fare',
    },
    {
      title: 'Remaining Seats',
      dataIndex: 'remaining_seats',
      key: 'remaining_seats',
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
    },
  ];

  return (
    <div>
      <h2>Flight Information</h2>
      <Table columns={columns} dataSource={flightInfo} />
    </div>
  );
};

export default Page2;
