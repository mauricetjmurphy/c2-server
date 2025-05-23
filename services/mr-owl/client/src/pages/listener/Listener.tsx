import axios from 'axios'
import { useEffect, useState } from 'react'
import List from '../../components/widgets/list/List'

const Listener = () => {
    const [rows, setRows] = useState([])

    useEffect(() => {
        axios
            .get('/mocks/getListeners.json')
            .then((resp) => {
                const data = resp.data.map((listener: any) => ({
                    ...listener,
                    devices: listener.devices
                        .map((device: any) => device.name)
                        .join(', '),
                }))
                setRows(data)
            })
            .catch((error) => {
                throw error
            })
    }, [])

    return (
        <List
            groupName="Listener"
            rows={rows}
            showButtonAdd={true}
            description="List of the available devices"
        />
    )
}

export default Listener
