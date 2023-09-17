import PropTypes from "prop-types";
import RoomRender from "./RoomRender";

function TableRender({ table, pos, it}) {
  return (
    <div className="flex space-x-2 rounded-12 p-1 bg-blue-">
      {table.map((e, i) => (
        <RoomRender state={e} current={pos === i} key={i} onClick={() => {
            typeof it === "function" && it(i, e);
        }}/>
      ))}
    </div>
  );
}

TableRender.propTypes = {
  table: PropTypes.object.isRequired,
  pos: PropTypes.number.isRequired,
  it: PropTypes.func,
};

export default TableRender;
