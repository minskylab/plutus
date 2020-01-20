import React from "react";
import {useRoutes, useRedirect} from 'hookrouter';

import HomePage from "./pages/home";
import AuthPage from "./pages/auth";
import SalesPage from "./pages/sales";
import DiscountsPage from "./pages/discounts";
import SettingsPage from "./pages/settings";

const routes = {
    "/": () => <HomePage/>,
    "/auth": () => <AuthPage/>,
    "/sales": () => <SalesPage/>,
    "/discounts": () => <DiscountsPage/>,
    "/settings": () => <SettingsPage/>,
};

const App = () => {
    useRedirect("/generic", "/");
    const currentPage = useRoutes(routes);

    return currentPage || <div> Page Not Found </div>;
};

export default App;
