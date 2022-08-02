<script>
    import {store} from './store.js';
    
    let username = '';
    export let value = undefined; 
    let show = value !== undefined
    let ref = null;
    
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