<script>
    import {store} from './store.js';
    import {onMount} from "svelte"
    
    let username = '';
    export let value = undefined; 
    let show = value !== undefined
    let ref = null;
    

    onMount(()=>{

        fetch("/api/user").then(data=>data.json()).then(js=>{
            username = js.name
            if(js.first_name ){
                value = js.first_name
            }else{
                value = js.name

            }

        })


    })



    function handleSubmit(e){
        e.preventDefault();
        if(username===""){
            return;
        }
        value = username;
        store.set({username: value});
        show = false;
    }    

    function handleKeyPress(e){
        if(e.key == "Enter"){
            handleSubmit(e);
        }
    }

    function resetUsername(){
        show = true
  
      
    }

    $: {
        if(show && ref){
            ref.focus()
        }
    }


</script>

{#if show}
<form on:submit={handleSubmit}>
    <input
        bind:this={ref}
        type="text"
        placeholder="Name"
        bind:value={username}
        id="name"
        on:keyup={handleKeyPress}

    />
    
</form>
{:else}
<p on:click={resetUsername}>{value}</p>
{/if}