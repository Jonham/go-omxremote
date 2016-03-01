var Controls=React.createClass({displayName:"Controls",handleClick:function(e){e.preventDefault();AjaxRequest(e.target.href,"POST")},render:function(){var backLink="/file/"+this.props.hash+"/backward",forwardLink="/file/"+this.props.hash+"/forward",pauseLink="/file/"+this.props.hash+"/pause";return React.createElement("div",{className:"controls"},React.createElement(FullAction,{Action:"Start",URL:"start",Hash:this.props.hash}),React.createElement("div",{className:"row"},React.createElement("div",{className:"one-third column"},React.createElement("a",{onClick:this.handleClick,className:"button u-full-width back",href:backLink},React.createElement("i",{className:"icon-back"})," Back")),React.createElement("div",{className:"one-third column"},React.createElement("a",{onClick:this.handleClick,className:"button u-full-width pause",href:pauseLink},React.createElement("i",{className:"icon-pause"})," Pause")),React.createElement("div",{className:"one-third column"},React.createElement("a",{onClick:this.handleClick,className:"button u-full-width forward",href:forwardLink},"Forward ",React.createElement("i",{className:"icon-forward"})))),React.createElement(FullAction,{Action:"Toggle Subs",URL:"subs",Hash:this.props.hash}),React.createElement(FullAction,{Action:"Stop",URL:"stop",Hash:this.props.hash}))}});
var FullAction=React.createClass({displayName:"FullAction",handleClick:function(e){e.preventDefault();AjaxRequest(e.target.href,"POST")},render:function(){var link="/file/"+this.props.Hash+"/"+this.props.URL,buttonClass="button u-full-width "+this.props.URL,iconClass="icon-"+this.props.URL;return React.createElement("div",{className:"row"},React.createElement("a",{className:buttonClass,onClick:this.handleClick,href:link},React.createElement("i",{className:iconClass}),this.props.Action))}});
var SearchBar=React.createClass({displayName:"SearchBar",handleChange:function(event){this.props.onUserInput(event.target.value)},render:function(){return React.createElement("div",{className:"search-container"},React.createElement("input",{type:"text",placeholder:"search",className:"u-full-width",value:this.props.searchText,ref:"filterTextInput",onChange:this.handleChange}))}});
var Video=React.createClass({displayName:"Video",getInitialState:function(){return {showResults:false}},onClick:function(e){e.preventDefault();if(this.state.showResults){this.setState({showResults:false})}else {this.setState({showResults:true})}},render:function(){var controls;if(this.state.showResults){controls=React.createElement(Controls,{hash:this.props.hash})}else {controls=null}return React.createElement("div",{key:this.props.hash,className:"video"},React.createElement("h5",null,React.createElement("a",{href:"#",onClick:this.onClick},this.props.file)),controls)}});
var Videos=React.createClass({displayName:"Videos",componentDidMount:function(){var that=this;AjaxRequest("/files","GET",function(response){that.setState({"Files":JSON.parse(response),"SearchText":that.state.SearchText})})},getInitialState:function(){return {"Files":[],"SearchText":""}},handleUserInput:function(filterText){this.setState({"Files":this.state.Files,"SearchText":filterText})},render:function(){var results=this.state.Files;return React.createElement("div",{className:"row"},React.createElement("h3",null,"omxremote"),React.createElement(SearchBar,{searchText:this.state.SearchText,onUserInput:this.handleUserInput}),results.map(function(result){if(fuzzy.test(this.SearchText,result.file)){return React.createElement(Video,{hash:result.hash,file:result.file})}},this.state))}});var AjaxRequest=function(url,method,callback){var xmlhttp=new XMLHttpRequest;xmlhttp.onreadystatechange=function(){if(this.readyState===4&&this.status==200){if(callback!==undefined)callback(this.responseText)}};xmlhttp.open(method,url,true);xmlhttp.send()};
