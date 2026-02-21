<script lang="ts">
    import Scan from "./Scan.svelte";
    import { StartScan } from '../../wailsjs/go/main/App';
    import { StartQuickScan } from '../../wailsjs/go/main/App';
    import Main from "./Main.svelte";

    let scanningStart: boolean = $state(false);
    let skipSetup: boolean = $state(false)

    function startQuickScan() {

        skipSetup = true;
        StartQuickScan()
    }

    function startScan() {
        // call go function to start scan function (savesearch)

        scanningStart = true;
        StartScan()
        
    }

    

</script>



{#if scanningStart === false && skipSetup === false}
    <div class="h-screen w-full">
        <div class="h-screen grid grid-cols-4 grid-rows-7 gap-4">
            <div class="col-span-2 col-start-2 row-start-2">
                <h1 class="relative flex items-center text-5xl font-medium text-[#D7D6FC] font-heebo">
                    First Time Setup
                </h1>
            </div>
            <div class="col-span-2 col-start-2 row-start-3 text-xl font-medium text-[#D7D6FC] font-heebo">
                <p>
                    To complete your first time set up, RetroSync will need to scan your device for compatible game saves.
                </p>
            </div>
            <div class="col-span-2 col-start-2 row-start-4 text-xl font-medium text-[#D7D6FC] font-heebo">
                <p>
                    Please note that this may take a few minutes. If your device has no game saves, you can choose to skip this step.
                </p>
            </div>
            <div class="col-start-2 row-start-6 grid place-items-center text-xl font-medium text-[#D7D6FC] font-heebo">
                <button onclick={startQuickScan} class="rounded-xl border border-(--glass-border) px-16 py-3 font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">Skip</button>
            </div>
            <div class="col-start-3 row-start-6 grid place-items-center text-xl font-medium text-[#D7D6FC] font-heebo">
                <button onclick={startScan} class="rounded-xl border border-(--glass-border) px-16 py-3 font-semibold shadow-lg/65 inset-shadow-sm inset-shadow-yellow/5 backdrop-blur-sm bg-(--glass-bg) inset-shadow-sm -fit cursor-pointer [&:hover]:scale-102 transition duration-10 active:inset-shadow-black/100" style="--bg: color-mix(in oklab, black 20%, transparent)">Next</button>
            </div>
        </div>
    </div>
{:else if scanningStart === true && skipSetup === false}
    <Scan />
{:else if skipSetup === true}
    <Main />
{/if}




<style>

@import "tailwindcss";

button {
  --glass-bg: color-mix(in oklab, var(--bg) 80%, transparent);
  --glass-border: color-mix(in oklab, var(--glass-bg) 80%, rgb(255, 253, 120));
}

</style>