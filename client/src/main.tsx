
import { ChakraProvider,createSystem,defaultConfig,defineConfig,mergeConfigs } from "@chakra-ui/react"

import React from "react"
import ReactDOM from "react-dom/client"
import App from "./App"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const theme = defineConfig({
        theme: {
            recipes: {
                Button: {
                    variants: {
                        variant: {
                            custom: {
                                borderRadius: "full",
                                bg: "black",
                                color: "black",
                                textTransform: "uppercase",
                            },
                        },
                    },
                },
            },
        }
    }
);

// Extends default theme
const config = mergeConfigs(defaultConfig, theme);
const system = createSystem(config);

export default system;

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
        <ChakraProvider value={system}>
            <App />
        </ChakraProvider>
    </QueryClientProvider>
  </React.StrictMode>,
)