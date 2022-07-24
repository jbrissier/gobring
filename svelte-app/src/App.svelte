<script>
	import Brings from "./brings.svelte";
	import { onMount } from "svelte";
	import _ from 'lodash'

	let username;

	onMount(()=>{
		username = window.localStorage.getItem("username")
			
	})

	$: {
		if(!_.isNil(username)) {
	
			window.localStorage.setItem("username", username);
		}
	}
</script>



<p>Hallo {username} </p>
<input type="text" placeholder="Username" bind:value={username} id="username">
<hr/>
<Brings />