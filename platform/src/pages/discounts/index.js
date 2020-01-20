import React from "react";
import GenericTemplate from "../../components/templates/generic";
import {usePath} from 'hookrouter';


const DiscountsPage = (props) => {
    const path = usePath();
    console.log(path);

    return <div>
        <GenericTemplate title={"Discounts"} selected={"discounts"}>
            <div>
                Hello World Discounts
            </div>
        </GenericTemplate>
    </div>
};

export default DiscountsPage;