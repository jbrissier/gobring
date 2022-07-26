<script>
    import {useNavigate} from 'svelte-navigator'
    const navigate = useNavigate();
    export let username;
    function handleSubmit() {
        const from = "/brings";
        navigate(from, { replace: true });
    }
</script>

<form on:submit={handleSubmit}>
    <p>Hallo {username}</p>
    <input
        type="text"
        placeholder="Username"
        bind:value={username}
        id="username"
    />
</form>
