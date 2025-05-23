import {
    Box,
    Button,
    Table,
    TableCaption,
    TableContainer,
    Tbody,
    Td,
    Text,
    Th,
    Thead,
    Tr,
} from '@chakra-ui/react'
import { nanoid } from 'nanoid'
import { IconPlus } from '@tabler/icons-react'

interface ListProps {
    groupName: string
    description: string
    showButtonAdd: boolean
    rows: { [key: string]: any }[]
}

const List = ({
    rows,
    groupName,
    showButtonAdd = false,
    description,
}: ListProps) => {
    if (rows.length > 0) {
        return (
            <>
                {showButtonAdd && (
                    <Box pb={3}>
                        <Button
                            variant="outline"
                            borderColor={'brand.purple'}
                            color={'brand.white'}
                            fontSize={'14px'}
                            leftIcon={<IconPlus size={'20px'} />}
                            sx={{
                                ':hover': {
                                    bg: '#00000050',
                                },
                            }}
                        >
                            {`Add ${groupName}`}
                        </Button>
                    </Box>
                )}

                <TableContainer
                    boxShadow="0 10px 15px rgba(0, 0, 0, 0.1), 0 4px 6px rgba(0, 0, 0, 0.05)"
                    borderRadius="md"
                    p={4}
                >
                    <Table variant="unstyled" colorScheme="customGreen">
                        <TableCaption>
                            <Text
                                color={'brand.grey_light'}
                                p={'8px'}
                                fontSize={'13px'}
                            >
                                {description}
                            </Text>
                        </TableCaption>
                        <Thead>
                            <Tr color={'brand.purple'}>
                                {Object.keys(rows[0])?.map((key) => (
                                    <Th key={nanoid()}>{key}</Th>
                                ))}
                                <Th textAlign={'right'}>Action</Th>
                            </Tr>
                        </Thead>
                        <Tbody
                            borderTop={'1px solid'}
                            borderColor={'brand.grey_light'}
                        >
                            {rows?.map((row) => (
                                <Tr
                                    key={nanoid()}
                                    borderBottom={'1px solid'}
                                    borderColor={'brand.grey_light'}
                                >
                                    {Object.keys(row).map((key) => (
                                        <Td
                                            key={nanoid()}
                                            fontFamily={'Roboto'}
                                            fontSize={'14px'}
                                            color={'brand.white'}
                                        >
                                            <Text pl={1}>{row[key]}</Text>
                                        </Td>
                                    ))}
                                    <Td textAlign={'right'}>
                                        <Button
                                            variant={'solid'}
                                            bg={'brand.purple'}
                                            border={'1px solid'}
                                            color={'brand.white'}
                                            borderColor={'brand.purple'}
                                            size={'sm'}
                                            onClick={() => {}}
                                            sx={{
                                                ':hover': {
                                                    bg: '#00000050',
                                                    borderColor: 'brand.purple',
                                                    color: 'brand.white',
                                                },
                                            }}
                                        >
                                            <Text fontSize={'12px'}>View</Text>
                                        </Button>
                                    </Td>
                                </Tr>
                            ))}
                        </Tbody>
                    </Table>
                </TableContainer>
            </>
        )
    }

    return null
}

export default List
