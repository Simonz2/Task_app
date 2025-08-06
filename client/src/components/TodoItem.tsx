import {Badge, Box,Flex,Spinner,Text} from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import type { Todo } from "./TodoList";
import { useMutation, useQueryClient } from "@tanstack/react-query";

import { BASE_URL } from "@/App";

const TodoItem = ({todo}:{todo:Todo})=>{
    console.log(todo);
const queryClient = useQueryClient();
const { mutate: updateTodo, isPending: isUpdating } = useMutation({
		mutationKey: ["updateTodo"],
		mutationFn: async () => {
            //console.log("Patching todo id: ", todo._id)
			if (todo.completed) return alert("Todo is already completed");
			
				const res = await fetch(BASE_URL + `/todos/${todo._id}`, {
					method: "PATCH",
				});
                
				const data = await res.json();
				if (!res.ok) {
					throw new Error(data.error || "Something went wrong");
				}
				return data;
			
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["todos"] });
		},
        onError:(err)=>{
            console.error(err);
            alert("something went wrong with updating todo");
        }
	});
//delete method
   const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
		mutationKey: ["deleteTodo"],
		mutationFn: async () => {
            //console.log("Patching todo id: ", todo._id)
			
				const res = await fetch(BASE_URL + `/todos/${todo._id}`, {
					method: "DELETE",
				});
                
				const data = await res.json();
				if (!res.ok) {
					throw new Error(data.error || "Failed to delete todo");
				}
				return data;
		
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["todos"] });//refresh the
		},
        onError:(err)=>{
            console.error(err);
            alert("something went wrong with deleting todo");
        },
	});
    return (
        <Flex gap={2} alignItems={"center"}>
            <Flex
                flex={1}
                alignItems={"center"}
                border={"1px"}
                borderColor={"gray.600"}
                p={2}
                borderRadius={"lg"}
                justifyContent={"space-between"}
                >
                < Text
                color={todo.completed ? "black.300":"violet"}
                    textDecoration={todo.completed ? "line-through":"none"}>
                    {todo.body}
                </Text>
                    <Badge ml="1" colorScheme= {todo.completed ? "green" : "yellow" } >
                        {todo.completed ? "Done" : "In progress"}
                    </Badge>
            </Flex> 
            <Flex>
                <Flex gap={2} alignItems={"center"}>
                    <Box color={"green.500"} cursor={"pointer"} onClick={() => updateTodo()}>
                        {isUpdating ? <Spinner size="sm"/> : <FaCheckCircle size={20}/>}
                        
                    </Box>
                    <Box color={"red.500"} cursor={"pointer"} onClick={()=> deleteTodo()}>
                        {isDeleting ? <Spinner size="sm"/>:<MdDelete size={25}/>}
                    </Box>
                </Flex>
            </Flex>
        </Flex>
    );
};
export default TodoItem