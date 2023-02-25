import bg from "../img/bg/3915e675f7a86ce55ddcc538388e788a1548144025.jpeg";
import { useSelector } from "react-redux";
import { useEffect, useState } from "react";
import { ExpandMore, LocationOn, Person } from "@mui/icons-material";
import { Button } from "@mui/material";
import { StarRating } from "./accessories/StarRating";

const Profile = () => {
  const { user } = useSelector((state) => state.auth);
  const [userLoc, setUserLoc] = useState(null);

  useEffect(() => {
    fetch("https://json.geoiplookup.io", { method: "GET" })
      .then((res) => res.json())
      .then((data) => setUserLoc(data))
      .catch((err) => console.log(err));
  }, []);

  const handleMenu = (e, id) => {
    document.querySelector(".activity .tab-header li.active").classList.remove("active");
    e.currentTarget.classList.add("active");
    document.querySelectorAll(".activity .tab-content").forEach((tab) => {
      if (tab.id === id) {
        document.querySelector(".activity .tab-content.active").classList.remove("active");
        tab.classList.add("active");
      }
    });
  };

  return (
    <div className="container">
      <section className="profile" style={{ padding: "20px 0" }}>
        <div
          className="header-con"
          style={{
            background: `linear-gradient(to right, rgba(0, 0, 0, 0.5),rgba(0, 0, 0, 0.5)), url(${bg}) no-repeat`,
          }}
        >
          <div>
            <div className="header-avi">
              <h1 style={{ fontSize: "50px" }}>{user?.email.charAt(0).toUpperCase()}</h1>
            </div>
            <div style={{ fontWeight: "bolder", marginLeft: "5px" }}>
              <h3>{user?.name}</h3>
              <div style={{ display: "flex" }}>
                <div className="input-icon">
                  <LocationOn color="#fff" fontSize="small" />
                </div>
                <h4 style={{ textAlign: "center" }}>{userLoc?.region}</h4>
              </div>
            </div>
          </div>
          <div>
            <Button variant="contained" size="small">
              Edit Profile
            </Button>
          </div>
        </div>
        <div className="activity">
          <div className="tab-header">
            <h4 className="tab-title" style={{ fontWeight: 400, margin: "0 10px 10px" }}>
              ACTIVITY
            </h4>
            <ul>
              <li className="active" onClick={(e) => handleMenu(e, "reviews")}>
                Reviews
              </li>
              <li onClick={(e) => handleMenu(e, "create-business")}>Create Business Page</li>
            </ul>
          </div>
          <div className="tab-content-con">
            <div className="tab-content reviews active" id="reviews">
              <h4
                className="tab-title"
                style={{
                  color: "#000",
                  borderBottom: "none",
                  paddingLeft: 0,
                  paddingTop: 0,
                  marginBottom: "5px",
                }}
              >
                Your Reviews
              </h4>
              {/* <hr className="hr-xl" /> */}
              <div style={{ display: "block" }}>
                <div>
                  <div className="review">
                    <div className="review-dp">
                      <div style={{ borderRadius: "10px" }}>
                        {/* {review.user.image_url ? (
                        <img src={review.user.image_url} alt="" />
                      ) : (
                        <Person style={{ color: "#cfcfcf", fontSize: "30px" }} />
                      )} */}
                        <Person style={{ color: "#cfcfcf", fontSize: "30px" }} />
                      </div>
                      <div>
                        <p>McDonalds</p>
                        <p>Nugeria</p>
                        {/* <p>{review.user.name}</p>
                        <p>{review.time_created}</p> */}
                      </div>
                      <div className="input-icon" style={{ marginLeft: "auto", placeItems: "start" }}>
                        <ExpandMore />
                      </div>
                    </div>
                    <div>
                      <StarRating value={4.5} />
                      <p>4.5</p>
                    </div>
                    <p>lorem ipsum</p>
                  </div>
                </div>
              </div>
            </div>
            <div className="tab-content create-business" id="create-business">
              <div>
                <h1>fuck it</h1>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
};

export default Profile;
