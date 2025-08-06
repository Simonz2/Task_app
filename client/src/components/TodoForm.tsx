import {Button, Flex, Input, Spinner} from "@chakra-ui/react"

import { useEffect, useRef, useState } from "react"
import { IoMdAdd } from "react-icons/io"
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { BASE_URL } from "@/App";

const TodoForm = () => {
    const [newTodo, setNewTodo] = useState("")
    const queryClient=useQueryClient();
    const inputRef=useRef<HTMLInputElement|null>(null);
    useEffect(()=>{
        inputRef.current?.focus();},[]);

    const {mutate: createTodo,isPending} =useMutation({
        mutationKey:["createTodo"],
        mutationFn: async()=>{
            if(!newTodo.trim()){
                alert("Enter a task");
            }
            const res=await fetch(BASE_URL+"/todos",{
                method:"POST",
                headers:{"Content-Type":"application/json",},
                body:JSON.stringify({body:newTodo}),
            });
            const data = await res.json();
            
            if (!res.ok) {
                throw new Error(data.error||"failed to create a todo");
            }
            return data   
            },
        onSuccess:()=>{
            queryClient.invalidateQueries({queryKey:["todos"]})
            setNewTodo("");//clear input
        },
        onError:(err)=>{
            console.error(err);
            alert("something went wrong :(");
        },
    });
    const handleSubmit= (e: React.FormEvent)=>{
        e.preventDefault();
        createTodo();
    };

    return (
    <form onSubmit={handleSubmit}>
        <Flex gap={2}>
            <Input
                type="text"
                value={newTodo}
                onChange={(e)=>setNewTodo(e.target.value)}
                ref={inputRef}
                placeholder="Enter new Todo"
                />
                <Button
                    mx={2}
                    type="submit"
                    colorScheme="teal"
                    _active={{transform: "scale(.97)",}}>
                        {isPending ? <Spinner size = {"xs"}/>:<IoMdAdd size={30}/>}
                    </Button>
        </Flex>
    </form>
    );

};
export default TodoForm;