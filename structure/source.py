import os
import subprocess

def generate_parent_tree_structure(output_file="parent_structure.txt"):
    """
    Generate a directory tree structure for the parent directory using Windows 'tree' command

    Args:
        output_file (str): Name of the output file (default: "parent_structure.txt")
    """
    try:
        # Get parent directory
        parent_dir = os.path.dirname(os.getcwd())

        if not parent_dir:
            print("Already at root directory, no parent exists")
            return

        print(f"Generating directory tree for parent: {parent_dir}")

        # Run the tree command in parent directory
        result = subprocess.run(['tree', '/F', '/A'],
                                cwd=parent_dir,
                                stdout=subprocess.PIPE,
                                stderr=subprocess.PIPE,
                                text=True,
                                shell=True)

        if result.returncode == 0:
            with open(output_file, 'w', encoding='utf-8') as f:
                f.write(f"Parent directory: {parent_dir}\n\n")
                f.write(result.stdout)
            print(f"Parent directory tree saved to {output_file}")

            # Print first few lines
            with open(output_file, 'r', encoding='utf-8') as f:
                print("\nFirst few lines of the output:")
                for i, line in enumerate(f):
                    if i < 10:
                        print(line, end='')
                    else:
                        break
        else:
            print("Error generating tree structure:")
            print(result.stderr)

    except Exception as e:
        print(f"An error occurred: {str(e)}")

if __name__ == "__main__":
    generate_parent_tree_structure()