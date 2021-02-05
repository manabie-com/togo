import React, { useEffect, useState } from 'react'
import axios from 'axios'
import './LoginForm.css'
import { API_BASE_URL, ACCESS_TOKEN_NAME } from '../../constants/apiContants'
import { withRouter } from "react-router-dom"

function LoginForm(props) {
    const [state, setState] = useState({
        userID: "",
        password: "",
        successMessage: null
    })
    
    const handleChange = (e) => {
        const { id, value } = e.target
        setState(prevState => ({
            ...prevState,
            [id]: value
        }))
    }

    useEffect(() => {
        const token = localStorage.getItem(ACCESS_TOKEN_NAME)
        if (token){
            redirectToHome()
        }
    })

    const handleSubmitClick = (e) => {
        e.preventDefault()
        axios.get(API_BASE_URL + '/login', {
            params: {
                "user_id": state.userID,
                "password": state.password,
            }
        })
            .then(resp => {
                if (resp.data && resp.data.data) {
                    setState(prevState => ({
                        ...prevState,
                        'successMessage': 'Login successful. Redirecting to home page..'
                    }))
                    localStorage.setItem(ACCESS_TOKEN_NAME, resp.data.data)
                    redirectToHome()

                    props.showError(null)
                } else {
                    props.showError("Unknow")
                }
            })
            .catch(err => {
                if (err.response && err.response.data && err.response.data.error){
                    props.showError(err.response.data.error)
                }else{
                    props.showError("Unknow")
                }
            })
    }
    const redirectToHome = () => {
        props.updateTitle('Home')
        props.history.push('/home')
    }

    return (
        <div className="card col-12 col-lg-4 login-card mt-2 hv-center">
            <form>
                <div className="form-group text-left">
                    <label htmlFor="exampleInputEmail1">User ID</label>
                    <input type="email"
                        className="form-control"
                        id="userID"
                        aria-describedby="emailHelp"
                        placeholder="Enter userID"
                        value={state.userID}
                        onChange={handleChange}
                    />
                </div>
                <div className="form-group text-left">
                    <label htmlFor="exampleInputPassword1">Password</label>
                    <input type="password"
                        className="form-control"
                        id="password"
                        placeholder="Password"
                        value={state.password}
                        onChange={handleChange}
                    />
                </div>
                <div className="form-check">
                </div>
                <button
                    type="submit"
                    className="btn btn-primary"
                    onClick={handleSubmitClick}
                >Submit</button>
            </form>
        </div>
    )
}

export default withRouter(LoginForm)