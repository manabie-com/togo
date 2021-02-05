import React, { useEffect, useState } from 'react'
import './Home.css'
import axios from 'axios'
import { API_BASE_URL, ACCESS_TOKEN_NAME, CREATE_DATE_NAME } from '../../constants/apiContants'
import { withRouter } from "react-router-dom"


function Home(props) {
    const [state, setState] = useState({
        listTask: [],
        content: 'first content',
        createDate: localStorage.getItem(CREATE_DATE_NAME) || '2020-06-29'
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
        if (!token) {
            redirectToLogin()
            return
        }

        getListTask(state.createDate)
    }, [props.state])

    const getListTask = (createDate) => {
        axios.get(API_BASE_URL + '/tasks', {
            params: { created_date: createDate },
            headers: { 'Authorization': localStorage.getItem(ACCESS_TOKEN_NAME) }
        })
            .then(resp => {
                if (resp.status === 200) {
                    setState(prevState => ({
                        ...prevState,
                        listTask: resp.data.data ? resp.data.data : []
                    }))
                } else {
                    redirectToLogin()
                }
            })
            .catch(() => {
                redirectToLogin()
            })
    }

    const addTask = (e) => {
        e.preventDefault()
        const payload = {
            "content": state.content
        }

        const config = {
            headers: {
                'Authorization': localStorage.getItem(ACCESS_TOKEN_NAME)
            }
        }

        axios.post(API_BASE_URL + '/tasks', payload, config)
            .then(resp => {
                if (resp.status === 200) {

                    const currentDate = getCurrentDate()

                    setState(prevState => ({
                        ...prevState,
                        createDate: currentDate
                    }))

                    getListTask(currentDate)
                } else {
                    redirectToLogin()
                }
            })
            .catch(err => {
                console.log(err.response)
                if (err.response && err.response.status === 401) {
                    redirectToLogin()
                }else{
                    props.showError(err.response.data.error || "Unknow")
                }
            })
    }

    const filter = (e) => {
        e.preventDefault()
        if (localStorage.getItem(CREATE_DATE_NAME) === state.createDate){
            return
        }
        localStorage.setItem(CREATE_DATE_NAME, state.createDate)
        getListTask(state.createDate)
    }

    const getCurrentDate = () => {
        const now = new Date()
        const date = now.getDate()
        const month = now.getMonth() + 1
        const year = now.getFullYear()
        return `${year}-${month < 10 ? `0${month}` : `${month}`}-${date}`
    }

    const redirectToLogin = () => {
        localStorage.removeItem(ACCESS_TOKEN_NAME)
        props.history.push('/login')
    }

    return (
        <div className="mt-2 wap">

            <form className="row g-3">
                <div className="col-auto">
                    <input type="text"
                        className="form-control"
                        id="content"
                        value={state.content}
                        onChange={handleChange}
                        placeholder="Content" />
                </div>
                <div className="col-auto">
                    <button type="button" onClick={addTask} className="btn  btn-success mb-3">Add task</button>
                </div>

                <div className="col-auto">
                    <input type="text"
                        className="form-control"
                        id="createDate"
                        value={state.createDate}
                        onChange={handleChange}
                        placeholder="Y-m-d" />
                </div>
                <div className="col-auto">
                    <button type="button" onClick={filter} className="btn  btn-success mb-3">Filter</button>
                </div>
            </form>

            <table className="table">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">ID</th>
                        <th scope="col">Content</th>
                        <th scope="col">User ID</th>
                        <th scope="col">Created Date</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        state.listTask.length > 0 && state.listTask.map((val, i) => {
                            return (
                                <tr key={i}>
                                    <th scope="row">{i + 1}</th>
                                    <td>{val.id}</td>
                                    <td>{val.content}</td>
                                    <td>{val.user_id}</td>
                                    <td>{val.created_date}</td>
                                </tr>
                            )
                        })
                    }
                </tbody>
            </table>
        </div>
    )
}

export default withRouter(Home)