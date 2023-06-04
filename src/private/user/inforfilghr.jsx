import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Table } from 'antd';

const Page2 = () => {
    const [flightInfo, setFlightInfo] = useState([]);

    useEffect(() => {
        const token = axios.defaults.headers.common['Authorization'];

        const getFlightInfo = async () => {
            try {
                const response = await axios.get('http://localhost:8080/user/info/flight');
                setFlightInfo(response.data.Info);
            } catch (error) {
                console.error('Lỗi trong quá trình lấy thông tin chuyến bay:', error);
            }
        };

        getFlightInfo();
    }, []);

    const columns = [
        {
            title: 'Mã chuyến bay',
            dataIndex: 'name_flight',
            key: 'name_flight',
        },
        {
            title: 'Hãng hàng không',
            dataIndex: 'name_airline',
            key: 'name_airline',
        },
        {
            title: 'Nơi đi',
            dataIndex: 'departure',
            key: 'departure',
        },
        {
            title: 'Nơi đến',
            dataIndex: 'destination',
            key: 'destination',
        },
        {
            title: 'Giờ khởi hành',
            dataIndex: 'departure_time',
            key: 'departure_time',
        },
        {
            title: 'Giờ đến',
            dataIndex: 'destination_time',
            key: 'destination_time',
        },
        {
            title: 'Loại vé',
            dataIndex: 'ticket_type',
            key: 'ticket_type',
        },
        {
            title: 'Giá vé',
            dataIndex: 'fare',
            key: 'fare',
        },
        {
            title: 'Số ghế còn lại',
            dataIndex: 'remaining_seats',
            key: 'remaining_seats',
        },
        {
            title: 'Trạng thái',
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