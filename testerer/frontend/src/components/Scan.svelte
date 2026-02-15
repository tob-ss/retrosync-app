<script lang="ts">
    import { CheckProgress } from '../../wailsjs/go/main/App'

    let progress: number = 0
    let barWidth = $state(0)

    function progressCheck() {
        CheckProgress().then(result => progress = result)
        barWidth = progress
    }

    setInterval(progressCheck, 500)

</script>

<div class="h-screen w-full">
        <div class="h-screen grid grid-cols-8 grid-rows-7 gap-4">
            <div class="col-span-4 col-start-3 row-start-2 self-center justify-self-center">
                <h1 class="relative flex items-center text-5xl font-medium text-[#D7D6FC] font-heebo">
                    First Time Setup
                </h1>
            </div>
            
            <div class="col-span-3 col-start-3 row-start-3 self-end text-xl font-medium text-[#D7D6FC] font-heebo">
                {#if barWidth === 0 }
                <p>Initialising Setup...</p>
                {:else if barWidth <= 35 && barWidth !== 0}
                <p>Searching Device for Game Saves...</p>
                {:else if barWidth <= 50 && barWidth > 35}
                <p>Verifying Found Game Saves...</p>
                {:else if barWidth < 100 && barWidth > 50}
                <p>Getting Metadata and Boxart...</p>
                {:else if barWidth === 100}
                <p>Setup Complete!</p>
                {/if}
            </div>
            <div class="col-start-6 row-start-3 justify-self-end self-end text-xl font-medium text-[#D7D6FC] font-heebo">
                <p>
                    {barWidth}%
                </p>
            </div>
            {#if barWidth >= 1}
            <div class="col-span-4 col-start-3 row-start-4 self-start">
                <div id="myProgress" class="rounded-xl border border-(--glass-border) py-3 font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit" style="--bg: color-mix(in oklab, white 20%, transparent); width: {barWidth}%">
                </div>
            </div>
            {/if}
            {#if barWidth === 100}
            <div class="col-start-6 row-start-6 col-span-2 grid place-items-center text-xl font-medium text-[#D7D6FC] font-heebo">
                <button class="rounded-xl border border-(--glass-border) px-16 py-3 font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">Next</button>
            </div>
            {/if}
        </div>
</div>

<style>

@import "tailwindcss";

#myProgress {
    width: 100%;
  --glass-bg: color-mix(in oklab, var(--bg) 80%, transparent);
  --glass-border: color-mix(in oklab, var(--glass-bg) 80%, rgb(255, 253, 120));
}

button {
  --glass-bg: color-mix(in oklab, var(--bg) 80%, transparent);
  --glass-border: color-mix(in oklab, var(--glass-bg) 80%, rgb(255, 253, 120));
}

</style>