import {
  Box,
  Flex,
  Text,
  Container,
  IconButton,
} from "@chakra-ui/react"

import {
    useColorModeValue,
    useColorMode,
} from "@/components/ui/color-mode"

import {LuMoon, LuSun} from "react-icons/lu";


export default function Navbar() {
    const {colorMode, toggleColorMode}=useColorMode();
    return (
        <Container maxW={"900px"}>
            <Box bg={useColorModeValue("gray.400","gray.700")} px={4} my={4} borderRadius={5}>
                <Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
                    {/*LEFT SIDE*/}
                    <Flex justifyContent={"center"} alignItems={"center"}
                        gap={2}
                        display={{base:"none",sm:"flex"}}
                        >
                            <img src="/react.png" alt="logo" width={50} height={50} />
                            <Text fontSize={"40"}>+</Text>
                            <img src="/golang.png" alt="logo" width={50} height={50} />
                            <Text fontSize={"40"}>=</Text>
                            <img src="/explode.png" alt="logo" width={50} height={50} />
                    </Flex>
                    {/*RIGHT SIDE*/}
                    <Flex alignItems={"center"} gap={3}>
                        <Text fontSize={"lg"} fontWeight={500}>
                            Daily Tasks
                        </Text>
                        {/*TOGGLE COLOR MODE*/}
                        <IconButton
                            onClick={toggleColorMode} variant="outline" size="sm">
                                {colorMode  === "light" ? <LuSun/> : <LuMoon/>}
                            </IconButton>
                    </Flex>
                </Flex>
            </Box>
        </Container>
    );
}