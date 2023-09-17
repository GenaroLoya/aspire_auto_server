import ax from "../libs/axios";

export const getSteps = async ({ queryKey }) => {
  const [_, { pos, table, dir }] = queryKey;

  console.log("getSteps", { pos, table, dir })
  return await ax
    .post("/steps", { pos, table, dir })
    .then((res) => res.data)
    .catch((err) => {
      if (err.response) {
        throw err.response.data;
      }
      throw err.request;
    });
};
