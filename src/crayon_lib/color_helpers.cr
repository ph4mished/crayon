#import strutils, strformat, os, tables, sequtils#, winim/lean
#import palette
#require "strutils"
require "./palette"
  


#a hex && rgb fallback will be added. where terminal doesn't support such, it reverts back to 256

module Crayon
    extend self

#===========================================
#  COLOR VALIDATION
#===========================================

  private def  valid_hex?(hex_code : String) : Bool
    #fg=#AABBCC
    begin
      return hex_code[4..-1].size == 6 &&  hex_code[4..-1].each_char.all? { |it| it.hex? || ('a'..'f').includes?(it.downcase)} && (hex_code.starts_with?("fg=") || hex_code.starts_with?("bg=")) #for rrggbb
    rescue
      return false
    end
  end


  private def valid_256_code?(palette_code : String) : Bool
    begin
      return palette_code[3..-1].to_i.in?(0..255) && (palette_code.starts_with?("fg=") || palette_code.starts_with?("bg="))
    rescue
      return false
    end
  end
    
  
  private def valid_rgb?(rgb_code : String) : Bool
    begin
      #fg=rgb(255,0,0)
      return rgb_code[3..6] == "rgb(" && rgb_code[-1] == ')' &&  rgb_code[7..-2].split(",").all? {|it| it.strip.each_char.all?  && it.strip.to_i.in?(0..255) } && (rgb_code.starts_with?("fg=") || rgb_code.starts_with?("bg="))
    rescue
      return false
    end
  end
  

  def true_color? : Bool
    return ENV["COLORTERM"]? == "truecolor" || ENV["COLORTERM"]? == "24bit"
  end

 
  #this function was made to validate words in []
  def supported_color?(input : String) : Bool
    return Crayon::Palette::COLOR_MAP.has_key?(input) || Crayon::Palette::RESET_MAP.has_key?(input) || Crayon::Palette::STYLE_MAP.has_key?(input) || valid_hex?(input) || valid_256_code?(input) || valid_rgb?(input)
  end


  #Get rgb values
  # fg=rgb(255,0,0)
  private def read_rgb(code : String) : Array(Int32) 
    rgb_values = [] of Int32
    code[7..-2].split(",").each do |num|
      rgb_values << num.strip.to_i
    end
    rgb_values
  end
    

  
  #======================================
  # COLOR PARSING
  #======================================

  private def parse_ansi(color_code : String, ansi_append : String) : String
    if color_code.starts_with?("bg=")
      return "\e[48;#{ansi_append}m"
    elsif color_code.starts_with?("fg=")
      return "\e[38;#{ansi_append}m"
    end
    ""
  end


  private def parse_rgb_to_ansi_code(rgb_code : String) : String
    if true_color? 
      rgb = read_rgb(rgb_code)
      return parse_ansi(rgb_code, "2;#{rgb[0]};#{rgb[1]};#{rgb[2]}")
    end
    ""
  end



  private def parse_hex_to_ansi_code(hex_code : String) : String
    if hex_code.size == 10
      #for #rrggbb
      if true_color?
        .
        r = hex_code[4..5].to_i(16)  #this is meant to extract RR
        g = hex_code[6..7].to_i(16)  # this is meant to extract GG
        b = hex_code[8..9].to_i(16)  #to extract BB
        return parse_ansi(hex_code, "2;#{r};#{g};#{b}")
      end
    end
    ""
  end

    #Note:
        #foreground colors use 38 && background colors use 48. the 2 is for truecolor support
      #so its \e[38;2;R;G;Bm || for background \e[48;2;R;G;Bm 
      #so the second row of number tells what color mode it is (2: rgb(24 bits), 245)
      # 2 is for truecolor supported numbers that is rgb && its 24 bits using a range of 0-255
      # 5 is for 256 palette(index 196) 
      # 256 palette support synax will be [fg=214] = foreground color && [bg=214] = background color
    
  private def parse_256_color_code(color_code : String) : String
    return parse_ansi(color_code, "5;#{color_code[3..-1]}")
  end


  def parse_color(color_code : String) : String
    #this function is meant to receive string like "bold" "fg=red" && other colors &&
    #convert them to their ansi codes
    if Crayon::Palette::COLOR_MAP.has_key?(color_code)
      return "\e[#{Crayon::Palette::COLOR_MAP[color_code]}m"

    elsif Crayon::Palette::STYLE_MAP.has_key?(color_code)
      return "\e[#{Crayon::Palette::STYLE_MAP[color_code]}m"

    elsif Crayon::Palette::RESET_MAP.has_key?(color_code)
      return "\e[#{Crayon::Palette::RESET_MAP[color_code]}m"

    elsif valid_256_code?(color_code)
      return parse_256_color_code(color_code)

    elsif valid_hex?(color_code)
      return parse_hex_to_ansi_code(color_code)

    elsif valid_rgb?(color_code)
      return parse_rgb_to_ansi_code(color_code)
      
    else
      ""
    end
  end
end
    
