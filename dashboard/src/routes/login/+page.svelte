<script lang="ts">
    import googleLogo from '$lib/assets/google-icon.png';
    import { browser } from '$app/environment';

    let loggedIn = $state(false);
 

    let email = $state('');
    let password = $state('');
    let errorMessage = $state('');

    function redirect() {
        if (browser) {
            window.location.href = '/dashboard/saves';
        }
        
    }

    function handleSubmit() {
        if (email === 'admin@test.com' && password === 'admin') {
            errorMessage = '';
            alert('Login successful!');
            loggedIn = true;
            email = '';
            password = '';
            
            
        } else {
            console.log("incorrect password")
            errorMessage = 'Invalid email address or password';
        }
    }

    


</script>

{#if loggedIn}
    {redirect()}
{:else}

<div class="h-screen grid grid-cols-7 grid-rows-10 gap-4 place-items-center">
    <div class="col-span-3 col-start-3 row-start-2 self-end">
        <h1 class="relative flex items-center text-center text-5xl font-medium text-[#D7D6FC] font-heebo">
            Login to RetroSync
        </h1>
    </div>
    
    <form class="w-full h-full col-span-3 row-span-5 col-start-3 row-start-3" on:submit|preventDefault={handleSubmit}>
        <div class="w-full h-full grid grid-cols-1 grid-rows-5 gap-4">
            <div class="w-full flex flex-col gap-4 row-start-2 text-xl font-medium text-[#D7D6FC] font-heebo self-end">
                <input class="p-4 rounded-xl border border-(--glass-border) font-semibold inset-shadow-black/100 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer" style="--bg: color-mix(in oklab, black 20%, transparent)" type="text" id="email" placeholder="Email Address" bind:value={email} />
                
                <input class="p-4 rounded-xl border border-(--glass-border) font-semibold inset-shadow-black/100 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer" style="--bg: color-mix(in oklab, black 20%, transparent)" type="password" id="password" placeholder="Password" bind:value={password} />
            </div>
            <div class="row-start-3 text-xl font-medium text-[#D7D6FC] font-heebo flex justify-between self-center">
                {#if errorMessage}
                    <p class="justify-self-start error">{errorMessage}</p>
                {/if}
                <p class="justify-self-end">
                    Forgot Password
                </p>
            </div>
            <div class="w-full h-full row-start-4 text-xl font-medium text-[#D7D6FC] font-heebo">
                <button class="w-full h-full rounded-xl border border-(--glass-border) font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">Login</button>
            </div>
        </div>
    </form>
    <div class="w-full h-full col-start-3 row-start-7 col-span-3 text-xl font-medium text-[#D7D6FC] font-heebo flex place-content-center items-center gap-4">
        <div class="w-1/4 h-1px border border-b-[#D7D6FC]-1"></div>
        <p class="justify-self-center">
            Or Login With
        </p>
        <div class="w-1/4 h-1px border border-b-[#D7D6FC]-1"></div>
    </div>
    <div class="w-full h-full col-start-3 row-start-8 col-span-3 text-xl font-medium text-[#D7D6FC] font-heebo">
        <button class="w-full h-full justify-items-center rounded-xl border border-(--glass-border) font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">
            <img src={googleLogo} alt="Google Login" width="40" height="40">
        </button>
    </div>
    <div class="w-full h-full col-start-3 row-start-9 col-span-3 text-xs md:text-sm lg:text-lg font-medium text-[#D7D6FC] font-heebo flex place-content-center items-center text-center gap-8">
        <p class="text-center">
            Don't have an account?
        </p>
        <button class="max-w-1/2 w-1/3 h-3/4 rounded-xl border border-(--glass-border) font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">
            Create one now
        </button>
    </div>
</div>

{/if}

<style>

@import "tailwindcss";

button {
  --glass-bg: color-mix(in oklab, var(--bg) 80%, transparent);
  --glass-border: color-mix(in oklab, var(--glass-bg) 80%, rgb(255, 253, 120));
}

input {
  --glass-bg: color-mix(in oklab, var(--bg) 80%, transparent);
  --glass-border: color-mix(in oklab, var(--glass-bg) 80%, rgb(255, 253, 120));
}


</style>