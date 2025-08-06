
import React from "react"
import ReactDOM from "react-dom/client"
import App from "./App"
import theme from "./theme"
import { ChakraProvider,ColorModeScript } from "@chakra-ui/react"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";


const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
        <ColorModeScript initialColorMode={theme.config.initialColorMode}/>
        <ChakraProvider theme={theme}>
                <App />
            
        </ChakraProvider>
    </QueryClientProvider>
  </React.StrictMode>,
)