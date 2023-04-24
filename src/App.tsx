import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Link from '@mui/material/Link';
import {
    Chip,
    CircularProgress,
    FormControlLabel,
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow, TextField
} from "@mui/material";
import { convertTypeAcquisitionFromJson } from 'typescript';

interface AppProps {
    endpoint: string
}

interface AppState {
    loading: boolean
    filter: string
    repos: Repo[]
}

interface Response {
    repos: Repo[]
}

interface Repo {
    org: string
    name: string
    description: string
    url: string
    labels: string[]
    contacts: {username: string, htmlurl: string}[]
    active: boolean
    archived: boolean
}

export default class App extends React.Component<AppProps, AppState> {
    xhr?: XMLHttpRequest
    constructor(props: AppProps) {
        super(props);
        this.state = {
            loading: true,
            filter: "",
            repos: []
        }
    }

    componentDidMount() {
        this.xhr = new XMLHttpRequest()
        const self = this
        this.xhr.addEventListener("readystatechange", function () {
            if (this.readyState === 4) {
                const response = JSON.parse(this.responseText) as Response
                self.setState({
                    loading: false,
                    repos: response.repos
                })
                self.xhr = undefined
            }
        })
        this.xhr.open("get", this.props.endpoint)
        this.xhr.send()
    }

    componentWillUnmount() {
        if (this.xhr) {
            this.xhr.abort()
        }
    }

    filter = (repo: Repo): boolean => {
        if (this.state.filter === "") {
            return true
        }
        const filterWords = this.state.filter.split(" ").filter((term) => term.length > 0)
        for (let filterWord of filterWords) {
            if (!repo.org.toLowerCase().includes(filterWord) &&
                !repo.name.toLowerCase().includes(filterWord) &&
                !repo.description.toLowerCase().includes(filterWord) &&
                repo.labels.filter((label) => label.toLowerCase().includes(filterWord)).length == 0 &&
                (!!repo.contacts ? repo.contacts.filter((contact) => contact.username.toLowerCase().includes(filterWord)).length == 0 : true)) {
                return false
            }
        }
        return true
    }

    render() {
        return this.state.loading?
            <Box sx={{ display: 'flex' }}>
                <CircularProgress />
            </Box>
            :
            <Box sx={{ display: 'flex', flexDirection: 'column' }}>
                <Box m={4}>
                    <Paper>
                        <Box p={4}>
                            <TextField
                                fullWidth
                                variant="standard"
                                label="Search"
                                value={this.state.filter}
                                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                                    this.setState({
                                        filter: event.target.value
                                    });
                                }}
                            />
                            Filter results: {this.state.repos.filter(this.filter).length}
                        </Box>
                    </Paper>
                </Box>
                <Box m={4}>
                Total repos: {this.state.repos.length}
                </Box>
                <Box m={4}>
                    <TableContainer component={Paper}>
                        <Table sx={{ minWidth: 650 }}>
                            <TableHead>
                                <TableRow>
                                    <TableCell>Organization</TableCell>
                                    <TableCell>Repository</TableCell>
                                    <TableCell>Description</TableCell>
                                    <TableCell>Labels</TableCell>
                                    <TableCell>Contacts</TableCell>
                                    <TableCell>Active</TableCell>
                                    <TableCell>Archived</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {this.state.repos.filter(this.filter).map((repo) => (
                                    <TableRow
                                        key={repo.org + " " + repo.name}
                                    >
                                        <TableCell>
                                            {repo.org}
                                        </TableCell>
                                        <TableCell>
                                            <a href={repo.url} target="_blank">{repo.name}</a>

                                        </TableCell>
                                        <TableCell>
                                            {repo.description}
                                        </TableCell>
                                        <TableCell>
                                            {repo.labels.map((label) => (
                                                <Chip key={repo.org + " " + repo.name + " " + label} label={label} color="primary" />
                                            ))}
                                        </TableCell>
                                        <TableCell>
                                            {!!repo.contacts && repo.contacts.map((contact,i) => <li key={i}><a href={contact.htmlurl} target="_blank">{contact.username}</a></li>)}
                                        </TableCell>
                                        <TableCell>
                                            {repo.active.toString()}
                                        </TableCell>
                                        <TableCell>
                                            {repo.archived.toString()}
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </Box>
            </Box>
    }
}
