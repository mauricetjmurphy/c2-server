import { useEffect, useState } from 'react'
import List from '../../components/widgets/list/List'
import axios from 'axios'

const Device = () => {
    const [rows, setRows] = useState([])

    useEffect(() => {
        axios
            .get('/mocks/getDevices.json')
            .then((resp) => setRows(resp.data))
            .catch((error) => {
                throw error
            })
    }, [])

    return (
        <List
            groupName="Device"
            rows={rows}
            showButtonAdd={true}
            description="List of the available devices"
        />
    )
}

export default Device
