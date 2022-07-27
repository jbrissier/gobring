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
<div class="container">


	<div class="plus">
		+
	</div>


	<div class={$brings.length === 0?"logo":"logo-top"}>
		
		<h1>GoBring</h1>
		<p>Bring good to your colleagues.</p>
		
	</div>
	
	<div class="login">
		
		<Login bind:value={username}/>
	</div>
	
	<div>
		<Brings/>
	</div>


</div>

<style>

	.container{
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
	
		height: 100%;
	}
	.plus{
		font-size: 4rem;
		margin: 0 1rem;
		cursor: pointer;
	}
	.logo{
		display: flex;
		flex-direction: column;
		margin: 0 2rem;

	}

	.logo-top{
		position: absolute;
		top: 1rem;
		flex-direction: column;
		margin: 0 2rem;
		left: 1rem;
	}

	.logo, .logo-top h1,p{
		margin: 0;
	}
	.login{
		position: fixed;
		top: 1rem;
		right: 1rem;
	}



</style>



