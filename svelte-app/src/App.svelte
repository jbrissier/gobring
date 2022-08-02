<script>
	import {store} from './store.js';
	import Login from "./Login.svelte";
	import Brings from "./brings.svelte";
	import { onMount } from "svelte";
	import _ from "lodash";
	import {brings} from './store.js'
	
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
<div class="login absolute right-3 top-3">
	
	<Login bind:value={username}/>
</div>
<div class=" bg-slate-400 flex flex-col w-screen h-screen p-10 justify-center items-center	">


	
	
	
	<div class="logo m-5">
		
		<h1 class="text-4xl text-slate-100">GoBring</h1>
		<p class="text-slate-900">Bring good to your colleagues.</p>
		
	</div>
	
	<div class="m-10">
		<Brings/>
	</div>
	
	

</div>

<style>

	*{
		font-family: 'Roboto Mono', monospace;
	}


</style>



