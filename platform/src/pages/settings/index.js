import React from "react";
import GenericTemplate from "../../components/templates/generic";
import {usePath} from 'hookrouter';


const SettingsPage = (props) => {
    const path = usePath();
    console.log(path);

    return <div>
        <GenericTemplate title={"Settings"} selected={"settings"}>
            <div>
                Hello World Settings
            </div>
        </GenericTemplate>
    </div>
};

export default SettingsPage;