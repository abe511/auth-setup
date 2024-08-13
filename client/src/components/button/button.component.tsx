import "./button.style.css";

type ButtonProps = {
  id?: string,
  type?: "submit" | "reset" | "button" | undefined,
  text?: string,
  className?: string,
}

const Button = ({id, type, text, className}: ButtonProps) => {
  return (
    <button id={id} type={type} className={`btn ${className}`}>{text}</button>
  )
};

export default Button;