<script>
	import {store} from './store.js';
	import Login from "./Login.svelte";
	import Brings from "./brings.svelte";
	import { onMount } from "svelte";
	import _ from "lodash";

	
	let username;



	onMount(() => {
		username = window.localStorage.getItem("username");
		store.set({username: username});
	});

	$: {
		if (!_.isNil(username)) {
			window.localStorage.setItem("username", username);
		}
	}
</script>

			<Login bind:value={username}/>
			<Brings />