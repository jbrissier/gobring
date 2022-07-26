<script>
    import {store} from './store.js';
    
    let username = '';
    export let value = undefined; 
    let ref = null;
    
    function handleSubmit(e){
        e.preventDefault();
        value = username;
        store.set({username: value});
    }    

    function handleKeyPress(e){
        if(e.key == "Enter"){
            handleSubmit(e);
        }
    }

    function resetUsername(){
        username = '';
        value = undefined;
        if(ref){
            ref.focus()
        }
    }


</script>

{#if value === undefined}
<form on:submit={handleSubmit}>
    <p>Hallo {username}</p>
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
<a  on:click={resetUsername}>{value}</a>
{/if}