<script>
	import { Router, Route, Link } from "svelte-navigator";
	import Login from "./Login.svelte";
	import Brings from "./brings.svelte";
	import { onMount } from "svelte";
	import _ from "lodash";

	
	let username;



	onMount(() => {
		username = window.localStorage.getItem("username");
	});

	$: {
		if (!_.isNil(username)) {
			window.localStorage.setItem("username", username);
		}
	}
</script>

<Router>
	<Link to="/bring">Brings</Link>
	<main>
		<Route path="/">
			<Login bind:username={username}/>
		</Route>
		<Route path="/brings">
			<Brings />
		</Route>
	</main>
</Router>
