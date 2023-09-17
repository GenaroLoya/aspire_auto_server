import PropTypes from "prop-types";
import classNames from "classnames";
import { STATES } from "../constants/main.aspire";

function RoomRender({ state, current, onClick }) {
  return (
    <div
      className={classNames(
        current ? "border-1 border-red-6" : "border-none",
        state === STATES.DIRTY ? "bg-amber-6" : "bg-amber-2",
        "relative rounded-12 p-1 w-8 h-8 flex justify-center items-center"
      )}
      onClick={onClick}
    >
      <span>{state}</span>
      {current && (
        <div className="absolute bg-orange-7 top-1 left-0 h-2 w-2 rounded-1"></div>
      )}
    </div>
  );
}

RoomRender.propTypes = {
  state: PropTypes.oneOf([STATES.CLEAN, STATES.DIRTY]).isRequired,
  current: PropTypes.bool.isRequired,
  onClick: PropTypes.func,
};

export default RoomRender;
