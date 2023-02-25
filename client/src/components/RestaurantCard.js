import { useNavigate } from "react-router-dom";
import { useSelector } from "react-redux";
import { useDispatch } from "react-redux";
import { TOGGLE_LOGIN } from "../reducers/auth";

import { StarRating } from "./accessories/StarRating";

const RestaurantCard = ({ restaurant, meals }) => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const { isAuth } = useSelector((state) => state.auth);

  const handleClick = () => {
    if (!isAuth) {
      return dispatch(TOGGLE_LOGIN());
    }
    navigate(`/restaurant/${restaurant.id}`);
  };
  return (
    <>
      <div className="best-food" onClick={handleClick}>
        <div>
          <img src={restaurant.image_url ? restaurant.image_url : meals[2].img} alt="" />
        </div>
        <div>
          <p>{restaurant.name}</p>
          <div className="rating">
            <StarRating value={restaurant.rating} />
            <p>{restaurant.rating}</p>
            <p>({restaurant.review_count} Reviews)</p>
          </div>
          <div className="categories">
            {restaurant.categories.map((x, idx) =>
              idx === restaurant.categories.length - 1 ? (
                <p key={idx}>{x.title}</p>
              ) : (
                <p key={idx}>{x.title},</p>
              )
            )}
          </div>
        </div>
      </div>
    </>
  );
};

export default RestaurantCard;
