"use client"
import { use, useState } from "react";
import Step1 from "./components/Step1";
import Step2 from "./components/Step2";
import { StepDoneIcon } from "./components/StepDoneIcon";
import { useRouter } from "next/navigation";

const Invest = () => {

    const [step, setStep] = useState(1)
    const [chain, setChain] = useState("")
    const [result, setResult] = useState<any[]>([])
    const router = useRouter()

    const handleSubmit = (asset: string) => {
        // Send the GET request
        let baseUrl = "http://localhost:3001/"
        baseUrl += chain + "/" + asset
        // Send the GET request using Fetch
        fetch(baseUrl)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                // Handle successful response
                console.log('Response:', data);
                setStep(3)
                setResult(data.data)
            })
            .catch(error => {
                // Handle error
                console.error('Error:', error);
            });
    }

    return (
        <div className="flex h-screen items-center justify-center p-24">
            <div className="flex justify-center items-start space-x-24">
                <ol className="col-span-2 relative text-gray-500 border-s border-gray-200 dark:border-gray-700 dark:text-gray-400">
                    <li onClick={() => setStep(1)} className="mb-10 ms-6">
                        <span className={`absolute flex items-center justify-center w-8 h-8 ${step > 1 ? 'bg-green-200' : 'bg-gray-100'} rounded-full -start-4 ring-4 ring-white dark:ring-gray-900 dark:bg-green-900`}>
                            {step > 1 ?
                                <StepDoneIcon />
                                :
                                <svg className="w-3.5 h-3.5 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 20">
                                    <path d="M16 1h-3.278A1.992 1.992 0 0 0 11 0H7a1.993 1.993 0 0 0-1.722 1H2a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2Zm-3 14H5a1 1 0 0 1 0-2h8a1 1 0 0 1 0 2Zm0-4H5a1 1 0 0 1 0-2h8a1 1 0 1 1 0 2Zm0-5H5a1 1 0 0 1 0-2h2V2h4v2h2a1 1 0 1 1 0 2Z" />
                                </svg>
                            }
                        </span>
                        <h3 className="font-medium leading-tight">Asset</h3>
                        <p className="text-sm">Where is your asset?</p>
                    </li>
                    <li className="mb-10 ms-6">
                        <span className={`absolute flex items-center justify-center w-8 h-8  ${step > 2 ? 'bg-green-200' : 'bg-gray-100'}  rounded-full -start-4 ring-4 ring-white dark:ring-gray-900 dark:bg-gray-700`}>
                            {step > 2 ?
                                <StepDoneIcon />
                                :
                                <svg className="w-3.5 h-3.5 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 20">
                                    <path d="M16 1h-3.278A1.992 1.992 0 0 0 11 0H7a1.993 1.993 0 0 0-1.722 1H2a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2ZM7 2h4v3H7V2Zm5.7 8.289-3.975 3.857a1 1 0 0 1-1.393 0L5.3 12.182a1.002 1.002 0 1 1 1.4-1.436l1.328 1.289 3.28-3.181a1 1 0 1 1 1.392 1.435Z" />
                                </svg>
                            }
                        </span>
                        <h3 className="font-medium leading-tight">Invesment</h3>
                        <p className="text-sm">Which type of protocol you want to invest?</p>
                    </li>
                </ol>
            </div>
            <div className="col-span-10 w-full h-full flex justify-center items-center">
                {step === 1 &&
                    <Step1 chainSelection={setChain} onStepDone={() => setStep((prev) => prev + 1)} />
                }
                {step === 2 &&
                    <Step2 assetChosen={handleSubmit} />
                }
                {step === 3 &&
                    <table>
                    <thead>
                      <tr className="text-black capitalize">
                        {/* Render table headers dynamically based on the first object in the data array */}
                        {result.length > 0 &&
                          Object.keys(result[0]).map((key) => (
                            <th key={key}>{key}</th>
                          ))}
                      </tr>
                    </thead>
                    <tbody>
                      {/* Render table rows dynamically */}
                      {result.length > 0 && result.map((row, index) => (
                        <tr className="text-black" key={index}>
                          {Object.values(row).map((value, index) => (
                            <td key={index}>{value as string | number}</td>
                          ))}
                        </tr>
                      ))}
                    </tbody>
                  </table>
                }

            </div>
        </div>
    )
}

export default Invest;