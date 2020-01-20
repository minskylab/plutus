import React, {useState, useEffect} from "react";
import { useLocalStorage } from "../../hooks";
import {
    Box,
    TextField,
    Container,
    Paper,
    Typography,
    Grid,
    Button,
    LinearProgress,
    Fade,
    Snackbar
} from "@material-ui/core";
import Image from 'material-ui-image';
import useFetch from 'use-http'
import {BASE_PLUTUS_URL} from "../../commons";

const AuthPage = () => {
    const [auth, setAuth] = useLocalStorage("auth", {});

    const [email, setEmail] = useState("");
    const [emailError, setEmailError] = useState("");

    const [password, setPassword] = useState("");
    const [passwordError, setPasswordError] = useState("");

    // const [loading, setLoading] = useState(false);
    // const [fatalError, setFatalError] = useState("");

    const [request, response] = useFetch(BASE_PLUTUS_URL);

    useEffect(() => {
        if (/\S+@\S+\.\S+/.test(email)) {
            setEmailError("");
        }else {
            setEmailError("Please type a valid email");
        }

        if (password.length < 6) {
            setPasswordError("Introduce 6 characters at least");
        }else {
            setPasswordError("");
        }

    }, [email, password]);


    const login = async () => {
        if (!!emailError || !!passwordError) {
            return;
        }

        await request.post("/auth/login", {
            "email": email,
            "password": password,
        });

        console.log(response.ok);
        // TODO... setAuth();
    };

    if (auth.ok) {
        window.location.replace("/");
        return <div>Redirecting</div>;
    }

    return (
        <React.Fragment>
            <Snackbar
                open={!!request.error}
                aria-describedby="message-id"
                message={<span id="message-id">{request.error?request.error.toString():""}</span>}
                autoHideDuration={2000}
                // onClose={() => ("")}
            />
            <Image
                src={"https://source.unsplash.com/collection/762960/1600x900"}
                aspectRatio={(window.innerWidth/window.innerHeight)}
                disableSpinner
                style={{zIndex: "-1", position:"absolute", width: "100%"}}
            />
            <Box display={"flex"} flexDirection={"column"} justifyContent={"center"} height={"100%"} zIndex={1}>
                <Container maxWidth={"xs"}>
                    <Paper elevation={3}>
                        <Box px={3} py={3} >
                            <Box pb={3}>
                                <Typography variant="h5" >
                                    Welcome to Plutus
                                </Typography>
                                <Typography variant={"body1"}>
                                    Please login with your credentials
                                </Typography>
                            </Box>

                            <Grid container>
                                <Grid item xs={12}>
                                    <Box py={1} display={"flex"} justifyContent={"center"}>
                                        <TextField
                                            value={email}
                                            error={!!emailError}
                                            label="Email"
                                            helperText={emailError}
                                            variant="outlined"
                                            onChange={(e)=>setEmail(e.target.value)}
                                        />
                                    </Box>
                                </Grid>
                                <Grid item xs={12}>
                                    <Box py={1} display={"flex"} justifyContent={"center"}>
                                        <TextField
                                            value={password}
                                            error={!!passwordError}
                                            label="Password"
                                            helperText={passwordError}
                                            variant="outlined"
                                            type="password"
                                            onChange={(e)=>setPassword(e.target.value)}
                                        />
                                    </Box>
                                </Grid>
                                <Grid item xs={12}>
                                    <Box py={1} display={"flex"} justifyContent={"center"}>
                                        <Button
                                            size="large"
                                            variant="contained"
                                            color="primary"
                                            disableElevation
                                            onClick={login}
                                            disabled={request.loading}
                                        >
                                            Login
                                        </Button>
                                    </Box>
                                </Grid>
                            </Grid>
                        </Box>
                        <Fade in={request.loading}>
                            <LinearProgress color="secondary" />
                        </Fade>
                    </Paper>
                </Container>
            </Box>
        </React.Fragment>

    );

};

export default AuthPage;
