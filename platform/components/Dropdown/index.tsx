
export type DropdownProps = {
    options: string[];
    text: string;
    setSelection: (selection:string)=>void
}

const Dropdown = ({ options, text, setSelection }: DropdownProps) => {
    const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedOption = event.target.value;
        setSelection(selectedOption);
    };

    return (
        <div className="py-10">
            <label className="form-control w-48">
                <div className="label">
                    <span className="text-xl text-black">{text}</span>
                </div>
                <select className="select bg-slate-700 select-bordered" onChange={handleChange}>
                    {options.map((option, index) => (
                        <option key={index} value={option}>{option}</option>
                    ))}
                </select>
            </label>
        </div>
    );
};

export default Dropdown;