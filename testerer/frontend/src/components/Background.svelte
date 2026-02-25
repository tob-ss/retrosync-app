<script lang="ts">
    import Welcome from "./Welcome.svelte"

</script>


<div class="dark relative flex flex-col h-screen items-center justify-center bg-[#2C2F48] transition-bg">
    <div class="absolute inset-0 overflow-hidden">
        <div class="dark jumbo absolute -inset-[10px] opacity-50">
            
        </div>
        
    </div>
    <Welcome />
</div>


<style>
@import url('https://fonts.googleapis.com/css2?family=Heebo:wght@100..900&display=swap');
@import "tailwindcss";


@theme {
    --font-heebo: "Heebo", sans-serif;
}

@keyframes jumbo {
    from {
        background-position: 50% 50%, 50% 50%;
    }
    to {
        background-position: 150% 50%, 150% 50%;
    }
}

.jumbo {
    --stripes: repeating-linear-gradient(
        100deg,
        #fff 0%,
        #fff 7%,
        transparent 10%,
        transparent 12%,
        #fff 16%
    );
    --stripesDark: repeating-linear-gradient(
        100deg,
        #000 0%,
        #000 7%,
        transparent 10%,
        transparent 12%,
        #000 16%
    );
    --rainbow: repeating-linear-gradient(
        100deg,
        #70B9EE 10%,
        #008FF5 15%,
        #70B9EE 20%,
        #008FF5 25%,
        #70B9EE 30%
    );
    background-image: var(--stripes), var(--rainbow);
    background-size: 300%, 200%;
    background-position: 50% 50%, 50% 50%;

    filter: invert(100%);

    mask-image: radial-gradient(ellipse at 100% 0%, black 40%, transparent 70%);

    pointer-events: none;
}

.jumbo::after {
    content: "";
    position: absolute;
    inset: 0;
    background-image: var(--stripes), var(--rainbow);
    background-size: 200%, 100%;
    animation: jumbo 600s linear infinite;
    background-attachment: fixed;
    mix-blend-mode: difference;
}

.dark .jumbo {
    background-image: var(--stripesDark), var(--rainbow);
    filter: blur(10px) opacity(50%) saturate(200%);
}
.dark .jumbo::after {
    background-image: var(--stripesDark), var(--rainbow);
}

</style>