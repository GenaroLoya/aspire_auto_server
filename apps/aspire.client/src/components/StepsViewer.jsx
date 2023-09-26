import { useQuery } from "react-query";
import { getSteps } from "../services/steps.service";
import TableRender from "./TableRender";
import { useEffect, useState } from "react";
import { DIRS, STATES } from "../constants/main.aspire";

function generateRandomObject() {
  // Generar un array aleatorio de longitud máxima 20 con valores 0 o 1
  const table = Array.from({ length: Math.floor(Math.random() * 20) + 1 }, () =>
    Math.random() < 0.5 ? 0 : 1
  );

  // Generar una posición aleatoria entre 0 y la longitud de la tabla generada
  const pos = Math.floor(Math.random() * table.length);

  // Generar una dirección aleatoria, 1 o -1
  const dir = Math.random() < 0.5 ? 1 : -1;

  // Crear y retornar el objeto con las características aleatorias
  const randomObject = { table, pos, dir };
  return randomObject;
}

function StepsViewer() {
  const [table, setTable] = useState([1, 1, 1, 0, 1]);
  const [pos, setPos] = useState(3);
  const [dir, setDir] = useState(DIRS.DIRR);

  const steps = useQuery(
    [
      "steps",
      {
        pos,
        table,
        dir,
      },
    ],
    getSteps
  );

  useEffect(() => {
    console.log(pos);
  }, [pos]);

  return (
    <div className="w-full flex justify-center">
      <div>
        {/**inputs for send data */}
        <div className="flex flex-col justify-center items-center space-x-2">
          <TableRender
            it={(pos, state) =>
              setTable((prev) => {
                const newTable = [...prev];
                newTable[pos] =
                  state === STATES.CLEAN ? STATES.DIRTY : STATES.CLEAN;
                return newTable;
              })
            }
            table={table}
            current={null}
          ></TableRender>
          <div className="relative bg-blueGray-6 p-1 flex h-20 items-center rounded-1 space-x-2 items-start">
            <button
              className="bg-green-500 h-9 w-9 hover:bg-green-700 text-white border-none py-1 px-2 rounded-1"
              onClick={() =>
                setTable((prev) => {
                  const newTable = [...prev];
                  newTable.push(STATES.CLEAN);

                  return newTable;
                })
              }
            >
              +
            </button>
            <button
              className="bg-red-500 h-9 w-9 hover:bg-red-700 border-none py-1 px-2 rounded-1"
              onClick={() =>
                setTable((prev) => {
                  if (prev.length >= 1) {
                    const newTable = [...prev];
                    newTable.pop();
                    if (pos >= newTable.length) setPos(newTable.length - 1);
                    return newTable;
                  }
                  return prev;
                })
              }
            >
              -
            </button>
            <div className="flex flex-col justify-center">
              <span>Initial Dir</span>
              <button
                className="bg-red-500 h-9 w-9 hover:bg-red-700 border-none py-1 px-2 rounded-1 ma space-y-.5"
                onClick={() => setDir((prev) => -prev)}
              >
                {dir === DIRS.DIRR ? "ᗧ" : "ᗤ"}
              </button>
            </div>
            <div className="flex flex-col justify-center">
              <span>Position</span>
              <input
                type="number"
                value={pos}
                onChange={(e) =>
                  setPos(() => {
                    console.log(e.target.value, pos, table.length);
                    const val = Number(e.target.value);
                    if (val >= table.length) return table.length - 1;
                    if (val < 0) return 0;
                    return val;
                  })
                }
                className=""
              />
            </div>
            <button
              className="absolute left-1/2 vertical-50% top-17 bg-yellow-7 h-9 hover:bg-green-700 text-white border-none py-1 px-2 rounded-1"
              onClick={() => {
                const randomObject = generateRandomObject();
                setTable(randomObject.table);
                setPos(randomObject.pos);
                setDir(randomObject.dir);
              }}
            >
              Generate random
            </button>
          </div>
        </div>

        <h1>Steps Viewer</h1>
        <div className="flex flex-col space-y-1 overflow-y-auto h-90">
          {steps?.data?.scene.map((e, i) => (
            <TableRender table={e.table} pos={e.pos}></TableRender>
          ))}
        </div>
      </div>
    </div>
  );
}

export default StepsViewer;
