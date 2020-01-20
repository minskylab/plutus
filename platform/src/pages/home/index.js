import React from "react";
import GenericTemplate from "../../components/templates/generic";
import {usePath} from 'hookrouter';


const HomePage = (props) => {
    const path = usePath();
    console.log(path);

    return <div>
        <GenericTemplate title={"Home"} selected={"home"}>
            <div>
                Hello World
            </div>
        </GenericTemplate>
    </div>
};

export default HomePage;