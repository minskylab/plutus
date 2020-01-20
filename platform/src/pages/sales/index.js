import React from "react";
import GenericTemplate from "../../components/templates/generic";
import {usePath} from 'hookrouter';


const SalesPage = (props) => {
    const path = usePath();
    console.log(path);

    return <div>
        <GenericTemplate title={"Sales"} selected={"sales"}>
            <div>
                Hello World Sales
            </div>
        </GenericTemplate>
    </div>
};

export default SalesPage;