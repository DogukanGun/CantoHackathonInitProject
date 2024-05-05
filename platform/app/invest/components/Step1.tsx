import CustomButton from "@/components/Button";
import { useEffect } from "react";
import { canto } from "viem/chains";
import { useSwitchChain } from "wagmi";
import ChainSelection from "./ChainSelection";

type Step1Param = {
    onStepDone: () => void
    chainSelection:(chain:string)=>void
}

const Step1 = ({ onStepDone,chainSelection }: Step1Param) => {
    const { switchChain } = useSwitchChain()
    useEffect(() => {
        switchChain({ chainId: canto.id })
    }, [])
    return (
        <div className="max-w-lg mx-auto">
            <ChainSelection chainSelection={chainSelection} />
            <CustomButton onClick={onStepDone} text="Next" />
        </div>
    )
}

export default Step1;