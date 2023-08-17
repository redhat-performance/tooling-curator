import * as React from 'react';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Link from '@mui/material/Link';
import Accordion from '@mui/material/Accordion';
import AccordionSummary from '@mui/material/AccordionSummary';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
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
  TableRow,
  TextField,
  IconButton,
} from "@mui/material";
import Fab from '@mui/material/Fab';
import ClearIcon from '@mui/icons-material/Clear';
import Tooltip from "@mui/material/Tooltip";
import { convertTypeAcquisitionFromJson } from 'typescript';
import SendIcon from "@mui/icons-material/Send";
import AccountCircleRoundedIcon from "@mui/icons-material/AccountCircleRounded";
import AndroidTwoToneIcon from "@mui/icons-material/AndroidTwoTone";
import ChatIcon from "@mui/icons-material/Chat";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import axios from "axios";
import { v4 as uuidv4 } from "uuid";
import ShadowMan from "./assets/redhat-shadowman.png";
import { width } from "@mui/system";
axios.defaults.headers.common["Access-Control-Allow-Origin"] = "*";

interface AppProps {
  endpoint: string
}

interface AppState {
  loading: boolean;
  filter: string;
  repos: Repo[];
  chatboxvisibility: boolean;
  messages: Message[];
  chatContent: string;
  uuid: string;
  isTooltipVisible: boolean;
}

interface Message {
  isBot: boolean;
  content: string;
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
  contacts: { username: string, htmlurl: string }[]
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
      repos: [],
      chatboxvisibility: false,
      messages: [
        {
          isBot: true,
          content:
            "Hello I'm Faro and I'm here to answer your questions on performance and scale tooling",
        },
      ],
      chatContent: "",
      uuid: uuidv4(),
      isTooltipVisible: false,
    };
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
      if (!repo.org.toLowerCase().includes(filterWord.toLowerCase()) &&
        !repo.name.toLowerCase().includes(filterWord.toLowerCase()) &&
        !repo.description.toLowerCase().includes(filterWord.toLowerCase()) &&
        repo.labels.filter((label) => label.toLowerCase().includes(filterWord.toLowerCase())).length == 0 &&
        (!!repo.contacts ? repo.contacts.filter((contact) => contact.username.toLowerCase().includes(filterWord.toLowerCase())).length == 0 : true)) {
        return false
      }
    }
    return true
  }

  render() {
    const reposByOrg: { [key: string]: Repo[] } = {};
    for (const repo of this.state.repos.filter(this.filter)) {
      if (!reposByOrg[repo.org]) {
        reposByOrg[repo.org] = [];
      }
      reposByOrg[repo.org].push(repo);
    }

    return this.state.loading ? (
      <Box sx={{ display: "flex" }}>
        <CircularProgress />
      </Box>
    ) : (
      <Box sx={{ display: "flex", flexDirection: "column" }}>
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
              <Tooltip title="Clear search" placement="bottom-start">
              <Fab  sx={{
                          position: 'absolute',
                          top: 60,
                          right: 60 }}
                      size='small'
                      color='warning'

                      onClick={() =>
                          this.setState({
                            filter: ""
                        })}>
                  <ClearIcon />
                </Fab>
              </Tooltip>
            </Box>
          </Paper>
        </Box>
        <Box m={4}>
          Total repos: {this.state.repos.length}
        </Box>
        <Box m={4}>
          {Object.entries(reposByOrg).map(([org, repos]) => (
            <Accordion key={org}>
              <AccordionSummary expandIcon={<ExpandMoreIcon />} >
                <Typography variant="h6">{org}</Typography>
              </AccordionSummary>
              <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }}>
                  <TableHead>
                    <TableRow>
                      <TableCell>Repository</TableCell>
                      <TableCell>Description</TableCell>
                      <TableCell>Labels</TableCell>
                      <TableCell>Contacts</TableCell>
                      <TableCell>Active</TableCell>
                      <TableCell>Archived</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {repos.map((repo) => (
                      <TableRow
                        key={repo.org + " " + repo.name}
                      >
                        <TableCell>
                          <a href={repo.url} target="_blank">{repo.name}</a>
                        </TableCell>
                        <TableCell>
                          {repo.description}
                        </TableCell>
                        <TableCell>
                          {repo.labels.map((label) => (
                            <Chip key={repo.org + " " + repo.name + " " + label} label={label} color="primary"
                                  onClick={() => {
                                    this.setState({
                                      filter: this.state.filter.toString() + " " + label
                                    })
                            }}/>
                          ))}
                        </TableCell>
                        <TableCell>
                          {!!repo.contacts && repo.contacts.map((contact, i) => <li key={i}><a href={contact.htmlurl} target="_blank">{contact.username}</a></li>)}
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
            </Accordion>
          ))}
        </Box>
        <div
          style={{
            textAlign: "right",
            position: "fixed",
            bottom: "20px",
            right: "20px",
            zIndex: "999",
          }}
        >
          <IconButton
            onClick={this.handleHelpclick}
            style={{
              color: "#EE0000",
              backgroundColor: "#c9c9c9",
              width: "60px",
              height: "60px",
            }}
          >
            <ChatIcon
              fontSize="large"
              style={{ width: "40px", height: "40px" }}
            />
          </IconButton>
        </div>
        {this.state.chatboxvisibility && (
          <div
            style={{
              height: "600px",
              width: "400px",
              position: "fixed",
              bottom: "40px",
              right: "60px",
              zIndex: "1000",
              borderRadius: "15px",
              boxShadow: "-2px 2px 10px rgba(0, 0, 0, 0.3)",
            }}
          >
            <div
              style={{
                height: "60px",
                width: "400px",
                backgroundColor: "#C1121F",
                borderRadius: "15px 15px 0px 0px",
                display: "flex",
                alignItems: "center",
                color: "white",
                fontSize: "20px",
                fontWeight: "bolder",
                padding: "10px",
              }}
            >
              <img
                src={ShadowMan}
                style={{ width: "40px", height: "40px" }}
              ></img>
              <span style={{ paddingLeft: "10px", fontSize: "25px" }}>
                <p>Faro</p>
              </span>
              <span style={{ paddingLeft: "10px", fontSize: "12px" }}>
                <p>(Experimental)</p>
              </span>
              <IconButton
                style={{ marginLeft: "auto" }}
                onClick={this.handleTooltipToggle}
              >
                <MoreVertIcon />
              </IconButton>
              {this.state.isTooltipVisible && (
                <div
                  className="tooltip-box"
                  style={{
                    zIndex: 2,
                    position: "absolute",
                    backgroundColor: "lightgrey",
                    marginTop: "-100px",
                    borderRadius: "5px",
                    width: "370px",
                    height: "30px",
                    fontSize: "12px",
                    padding: "3px",
                    border: "1px solid black",
                    color: "black",
                  }}
                >
                  To submit feedback please mail:{" "}
                  <a href={`mailto:smalleni@redhat.com`} color="red">
                    smalleni@redhat.com
                  </a>
                </div>
              )}
            </div>

            <div
              style={{
                backgroundColor: "#f2f2f2",
                height: "480px",
                overflowY: "auto",
                padding: "10px",
              }}
            >
              {this.state.messages.map((message, index) => (
                <div
                  id={index.toString()}
                  style={{
                    textAlign: message.isBot ? "left" : "right",
                    alignItems: "center",
                  }}
                >
                  {message.isBot ? (
                    <img
                      src={ShadowMan}
                      style={{ width: "30px", height: "30px" }}
                    ></img>
                  ) : null}
                  <div
                    style={{
                      maxWidth: "300px",
                      wordWrap: "break-word",
                      padding: "8px 12px",
                      backgroundColor: "lightgrey",
                      borderRadius: "10px",
                      display: "inline-block",
                      margin: "10px",
                      textAlign: "left",
                    }}
                  >
                    {message.content}
                  </div>
                  {!message.isBot ? (
                    <AccountCircleRoundedIcon style={{ width: "30px" }} />
                  ) : null}
                </div>
              ))}
            </div>
            <div
              style={{
                position: "absolute",
                bottom: "0px",
                height: "60px",
                width: "400px",
                backgroundColor: "#f2f2f2",
                borderRadius: "0px 0px 15px 15px",
                color: "white",
                borderTop: "3px solid black",
                alignItems: "center",
              }}
            >
              <TextField
                label="What's your question"
                variant="standard"
                autoComplete="off"
                style={{
                  width: "330px",
                  margin: "0px 0px 0px 10px",
                }}
                onKeyDown={this.handleKeyPress}
                value={this.state.chatContent}
                onChange={this.handleInputChange}
              />
              <IconButton
                color="primary"
                style={{
                  right: "10px",
                  margin: "10px",
                  backgroundColor: "lightgrey",
                  color: "#C1121F",
                }}
                onClick={this.handleSend}
              >
                <SendIcon />
              </IconButton>
            </div>
          </div>
        )}
      </Box>
    );
  }

  handleHelpclick = () => {
    this.setState({
      chatboxvisibility: !this.state.chatboxvisibility,
    });
  };

  handleSend = () => {
    if (this.state.chatContent === "") {
      alert("Empty input, ask your question");
      return;
    }
    if (!this.state.messages[this.state.messages.length - 1].isBot) {
      alert("Please wait until the bot replies!");
      return;
    }
    console.log(`New message incoming!`);
    this.setState({
      messages: [
        ...this.state.messages,
        { isBot: false, content: this.state.chatContent },
      ],
    });

    axios
      .get("https://faro-tooling-curator.com/ask", {
        headers: { uuid: this.state.uuid },
        params: { query: this.state.chatContent },
      })
      .then((res) => {
        console.log(res);
        this.setState({
          messages: [
            ...this.state.messages,
            { isBot: true, content: res.data.answer },
          ],
        });
      })
      .catch((err) => {
        console.log(err);
      });

    this.setState({
      chatContent: "",
    });
  };

  handleTooltipToggle = () => {
    this.setState({
      isTooltipVisible: !this.state.isTooltipVisible,
    });
  };

  handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    this.setState({
      chatContent: event.target.value,
    });
  };

  handleKeyPress = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      this.handleSend();
    }
  };
}
