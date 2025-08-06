import {Box,Container, VStack} from '@chakra-ui/react';
import Navbar from './components/Navbar';
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'

export const BASE_URL="http://localhost:5000/api"

function App() {
  return (
    <Box minH="100vh" bg="gray.100" _dark={{bg:"gray.900"}}>
      <Navbar />
      <Container>
        <VStack  spacing={6} align="stretch">
          
          <TodoForm/>
          <TodoList/>
        </VStack>
        
      </Container>
    </Box>
  )
}

export default App
