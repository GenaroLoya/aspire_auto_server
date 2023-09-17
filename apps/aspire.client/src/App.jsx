import { QueryClient, QueryClientProvider } from "react-query";
import StepsViewer from "./components/StepsViewer";
import { ReactQueryDevtools } from "react-query/devtools";

const queryClient = new QueryClient();

function App() {
  return (
    // Provide the client to your App
    <QueryClientProvider client={queryClient} contextSharing={true}>
      <div className="flex items-center justify-center h-screen @dark:bg-slate-8 @dark:text-neutral-3">
        <StepsViewer/>
      </div>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;
